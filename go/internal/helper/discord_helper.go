package helper

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/bwmarrin/discordgo"
)

func RemoveGlobalCommands(s *discordgo.Session) error {
	log.Sugar.Debug("Removing Global commands...")

	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Sugar.Error("Could not fetch registered commands")
		return err
	}

	for _, v := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Sugar.Errorf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}

	return nil
}

func IsInVoiceChannel(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	return s.VoiceConnections[i.GuildID] != nil
}

func GetUserVoiceChannel(s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {
	vs, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil {
		log.Sugar.Errorf("Failed to get voice state: %v", err)
		return "", err
	}

	return vs.ChannelID, nil
}

func PlayAudio(vc *discordgo.VoiceConnection, filePath string) error {
	buffer, err := loadSound(filePath)
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	vc.Speaking(true)

	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	vc.Speaking(false)

	time.Sleep(1 * time.Second)

	return nil
}

func loadSound(filePath string) ([][]byte, error) {
	var buffer = make([][]byte, 0)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return nil, err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return buffer, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}
