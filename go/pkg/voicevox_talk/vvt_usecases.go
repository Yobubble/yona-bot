package vvt

import (
	"fmt"
	"os"
	"time"

	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
)

type VVTUseCases struct {
	Vc *discordgo.VoiceConnection
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *VVTUseCases) createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
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

func (v *VVTUseCases) voiceRecording() ([]uint32, error) {
	files := make(map[uint32]media.Writer)
	for {
		select {
		case p := <-v.Vc.OpusRecv:
			file, ok := files[p.SSRC]
			if !ok {
				var err error
				file, err = oggwriter.New(fmt.Sprintf("./assets/audios/ssrcs/ogg/%d.ogg", p.SSRC), 48000, 2)
				if err != nil {
					log.Sugar.Error("Error creating audio file")
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
		case <-time.After(2 * time.Second): // stop receiving packets after silence of 2 seconds
			log.Sugar.Info("Silence detected...")

			log.Sugar.Debugf("There %d user(s) in the voice channel", len(files))

			ssrcs := make([]uint32, 0, len(files))
			for ssrc, file := range files {
				log.Sugar.Debugf("ssrc: %d", ssrc)
				ssrcs = append(ssrcs, ssrc)
				file.Close()
			}

			// NOTE: if the voiceConnection is closed then you have to re-join the channel to interact again
			// close(v.OpusRecv)
			// v.Close()
			return ssrcs, nil
		}
	}
}

func (v *VVTUseCases) createNewChatHistory(guildName string) error {
	filePath := fmt.Sprintf("./assets/chat_history/%s.txt", guildName)

	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}

	_, err = os.Create(filePath)
	if err != nil {
		return err
	}

	return nil
}
