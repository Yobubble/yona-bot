package vvt

import (
	"errors"
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
					return nil, fmt.Errorf("voicevox talk usecase: error creating audio file: %w", err)
				}
				files[p.SSRC] = file
			}

			// Construct pion RTP packet from DiscordGo's type.
			rtp := v.createPionRTPPacket(p)
			err := file.WriteRTP(rtp)
			if err != nil {
				return nil, fmt.Errorf("voicevox talk usecase: error writing to audio file: %w", err)
			}

		case <-time.After(2 * time.Second):
			log.Sugar.Info("Silence detected...")

			log.Sugar.Debugf("There is/are %d user(s) in the voice channel", len(files))
			if len(files) == 0 {
				return nil, errors.New("no user in the voice channel")
			}

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
	paths := []string{
		enum.SSRC_OGG.GetPath(),
		enum.SSRC_MP3.GetPath(),
		enum.VVE.GetPath(),
		enum.Audio.GetPath(),
	}

	for _, val := range paths {
		if err := os.MkdirAll(val, os.ModePerm); err != nil {
			return fmt.Errorf("voicevox talk usecase: error make %s directory: %w", val, err)
		}
	}

	return nil
}

func (v *vvtUseCase) post() error {
	paths := []string{
		enum.SSRC_OGG.GetPath(),
		enum.SSRC_MP3.GetPath(),
		enum.VVE.GetPath(),
		enum.Audio.GetPath(),
	}

	for _, val := range paths {
		if err := os.RemoveAll(val); err != nil {
			return fmt.Errorf("voicevox talk usecase: error remove %s directory: %w", val, err)
		}
	}

	return nil
}

func (v *vvtUseCase) processRecordAudio(ssrcStr string, guildName string) error {
	start := time.Now()

	// ogg -> mp3
	if err := v.ah.ConvertToMp3(enum.SSRC_OGG.GetFullPath(ssrcStr), enum.SSRC_MP3.GetFullPath(ssrcStr)); err != nil {
		return fmt.Errorf("voicevox talk usecase: error converting ogg to mp3: %w", err)
	}

	// STT
	log.Sugar.Debug("Speech To Text...")
	question, err := v.stt.AudioToText(enum.SSRC_MP3.GetFullPath(ssrcStr), enum.JP)
	if err != nil {
		return fmt.Errorf("voicevox talk usecase: error converting audio to text: %w", err)
	}

	// LM
	log.Sugar.Debug("LM Answering...")
	answer, err := v.lm.AskQuestion(guildName, question)
	if err != nil {
		return fmt.Errorf("voicevox talk usecase: error generating answer: %w", err)
	}

	if err := v.lm.UpdateChatHistory(guildName, question, answer); err != nil {
		return fmt.Errorf("voicevox talk usecase: error updating chat history: %w", err)
	}

	// TTS
	log.Sugar.Debug("Text To Speech...")
	err = v.tts.TextToSpeech(answer, enum.VVE.GetFullPath(ssrcStr))
	if err != nil {
		return fmt.Errorf("voicevox talk usecase: error converting text to audio: %w", err)
	}

	// mp3 -> dca
	err = v.ah.ConvertToDCA(enum.VVE.GetFullPath(ssrcStr), enum.Audio.GetFullPath(ssrcStr))
	if err != nil {
		return fmt.Errorf("voicevox talk usecase: error converting mp3 to dca: %w", err)
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
