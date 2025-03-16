package voicevoxTalk

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/Yobubble/yona-bot/config"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/oggwriter"
	"go.uber.org/zap"
)

func isInVoiceChannel(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	return s.VoiceConnections[i.GuildID] != nil
}

func getUserVoiceChannel(s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {
	vs, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil {
		config.Sugar.Warn("Failed to get voice state:", zap.Error(err))
		return "", err
	}
	return vs.ChannelID, nil
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version: 2,
			// Taken from Discord voice docs
			PayloadType:    0x78,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

func voiceRecording(v *discordgo.VoiceConnection) ([]uint32, error) {
	files := make(map[uint32]media.Writer)

	for {
		select {
		case p := <-v.OpusRecv:
			config.Sugar.Info("Voice packet receiving started")
			file, ok := files[p.SSRC]
			if !ok {
				var err error
				file, err = oggwriter.New(fmt.Sprintf("/go/voicevox_talk/audios/%d.ogg", p.SSRC), 48000, 2)
				if err != nil {
					config.Sugar.Warn("Error creating audio file")
					return nil, err
				}
				files[p.SSRC] = file
			}
			// Construct pion RTP packet from DiscordGo's type.
			rtp := createPionRTPPacket(p)
			err := file.WriteRTP(rtp)
			if err != nil {
				config.Sugar.Warn("Error writing to audio file")
				return nil, err
			}
		case <-time.After(2 * time.Second): // stop receiving packets after silence of 3 seconds
			config.Sugar.Info("Silence detected...")
			ssrcs := make([]uint32, len(files))
			for ssrc, file := range files {
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

func convertOggToMp3(fileName string) error {
	app := "ffmpeg"
	flag1 := "-i" // input - path to the file
	arg1 := fmt.Sprintf("audios/%s.ogg", fileName)
	flag2 := "-vn" // no video
	flag3 := "-ar" // audio rate - 48000
	arg3 := "48000"
	flag4 := "-ac" // audio channels - 2
	arg4 := "2"
	flag5 := "-b:a" // audio bitrate - 192k
	arg5 := "192k"
	arg6 := fmt.Sprintf("audios/%s.mp3", fileName)

	cmd := exec.Command(app, flag1, arg1, flag2, flag3, arg3, flag4, arg4, flag5, arg5, arg6)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
