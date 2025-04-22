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

type DiscordHelper struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

func (d *DiscordHelper) RemoveGlobalCommands() error {
	cmds, err := d.S.ApplicationCommands(d.S.State.User.ID, "")
	if err != nil {
		log.Sugar.Error("Could not fetch registered commands")
		return err
	}

	for _, v := range cmds {
		err := d.S.ApplicationCommandDelete(d.S.State.User.ID, "", v.ID)
		if err != nil {
			log.Sugar.Errorf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}

	log.Sugar.Debug("Successfully removed global commands...")
	return nil
}

func (d *DiscordHelper) GetGuildName() (string, error) {
	guild, err := d.S.Guild(d.I.GuildID)
	if err != nil {
		log.Sugar.Errorf("Failed to get guild name: %v", err)
		return "", err
	}

	return guild.Name, nil
}

func (d *DiscordHelper) IsBotInVoiceChannel() bool {
	return d.S.VoiceConnections[d.I.GuildID] != nil
}

func (d *DiscordHelper) GetVoiceConnection() *discordgo.VoiceConnection {
	return d.S.VoiceConnections[d.I.GuildID]
}

func (d *DiscordHelper) GetUserVoiceChannel() (string, error) {
	vs, err := d.S.State.VoiceState(d.I.GuildID, d.I.Member.User.ID)
	if err != nil {
		log.Sugar.Errorf("Failed to get voice state: %v", err)
		return "", err
	}

	return vs.ChannelID, nil
}

func (d *DiscordHelper) PlayAudio(vc *discordgo.VoiceConnection, filePath string) error {
	buffer, err := d.loadSound(filePath)
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

func (d *DiscordHelper) loadSound(filePath string) ([][]byte, error) {
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

func NewDiscordHelper(s *discordgo.Session, i *discordgo.InteractionCreate) *DiscordHelper {
	return &DiscordHelper{
		S: s,
		I: i,
	}
}
