package vvt

import (
	"fmt"
	"os"
	"time"

	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	lm "github.com/Yobubble/yona-bot/pkg/lm"
	"github.com/Yobubble/yona-bot/pkg/storage"
	"github.com/Yobubble/yona-bot/pkg/stt"
	"github.com/Yobubble/yona-bot/pkg/tts"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
)

type vvtUseCase struct {
	st  storage.Storage
	ah  *helper.AudioHelper
	lm  lm.LM
	tts tts.TTSModel
	stt stt.STTModel
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *vvtUseCase) createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version:        2,
			PayloadType:    0x78,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

func (v *vvtUseCase) voiceRecording(vc *discordgo.VoiceConnection) ([]uint32, error) {
	files := make(map[uint32]media.Writer)
	for {
		select {
		case p := <-vc.OpusRecv:
			file, ok := files[p.SSRC]
			if !ok {
				var err error
				fileName := fmt.Sprintf("%d", p.SSRC)
				filePath := enum.SSRC_OGG.GetFullPath(fileName)

				file, err = oggwriter.New(filePath, 48000, 2)
				if err != nil {
					log.Sugar.Errorf("Error creating audio file: %v", err)
					return nil, err
				}
				files[p.SSRC] = file
			}

			// Construct pion RTP packet from DiscordGo's type.
			rtp := v.createPionRTPPacket(p)
			err := file.WriteRTP(rtp)
			if err != nil {
				log.Sugar.Error("Error writing to audio file")
				return nil, err
			}

		case <-time.After(2 * time.Second):
			log.Sugar.Info("Silence detected...")

			log.Sugar.Debugf("There %d user(s) in the voice channel", len(files))

			ssrcs := make([]uint32, 0, len(files))
			for ssrc, file := range files {
				log.Sugar.Debugf("ssrc: %d", ssrc)
				ssrcs = append(ssrcs, ssrc)
				file.Close()
			}

			// IMPORTANT: if the voiceConnection is closed then you have to re-join the channel to interact again
			return ssrcs, nil
		}
	}
}

func (v *vvtUseCase) pre() error {
	if err := os.MkdirAll(enum.SSRC_OGG.GetPath(), os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(enum.SSRC_MP3.GetPath(), os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(enum.VVE.GetPath(), os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(enum.Audio.GetPath(), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (v *vvtUseCase) post() error {
	if err := os.RemoveAll(enum.SSRC_OGG.GetPath()); err != nil {
		return err
	}

	if err := os.RemoveAll(enum.SSRC_MP3.GetPath()); err != nil {
		return err
	}

	if err := os.RemoveAll(enum.VVE.GetPath()); err != nil {
		return err
	}

	if err := os.RemoveAll(enum.Audio.GetPath()); err != nil {
		return err
	}

	return nil
}

func (v *vvtUseCase) processRecordAudio(ssrcStr string, guildName string) error {
	start := time.Now()

	// ogg -> mp3
	if err := v.ah.ConvertToMp3(enum.SSRC_OGG.GetFullPath(ssrcStr), enum.SSRC_MP3.GetFullPath(ssrcStr)); err != nil {
		log.Sugar.Errorf("Error converting from ogg to mp3: %v", err)
		return err
	}

	// STT
	log.Sugar.Debug("Speech To Text...")
	question, err := v.stt.AudioToText(enum.SSRC_MP3.GetFullPath(ssrcStr), enum.JP)
	if err != nil {
		log.Sugar.Errorf("Error converting audio to text:%v", err)
		return err
	}

	// LM
	log.Sugar.Debug("LM Answering...")
	answer, err := v.lm.AskQuestion(guildName, question)
	if err != nil {
		log.Sugar.Errorf("Error generating answer: %v", err)
		return err
	}

	if err := v.lm.UpdateChatHistory(guildName, question, answer); err != nil {
		log.Sugar.Errorf("Error updating chat history: %v", err)
		return err
	}

	// TTS
	log.Sugar.Debug("Text To Speech...")
	err = v.tts.TextToSpeech(answer, enum.VVE.GetFullPath(ssrcStr))
	if err != nil {
		log.Sugar.Errorf("Error converting from text to audio: %v", err)
		return err
	}

	// mp3 -> dca
	err = v.ah.ConvertToDCA(enum.VVE.GetFullPath(ssrcStr), enum.Audio.GetFullPath(ssrcStr))
	if err != nil {
		log.Sugar.Errorf("Error converting from mp3 to dca: %v", err)
		return err
	}

	end := time.Now()
	log.Sugar.Debugf("Take %s seconds", end.Sub(start))

	return nil
}

func NewVVTUseCase(st storage.Storage, lm lm.LM, tts tts.TTSModel, stt stt.STTModel, ah *helper.AudioHelper) *vvtUseCase {
	return &vvtUseCase{
		st:  st,
		lm:  lm,
		tts: tts,
		stt: stt,
		ah:  ah,
	}
}
