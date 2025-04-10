package discordcmd

import (
	"time"

	"github.com/Yobubble/yona-bot/internal/log"
	voicevoxTalk "github.com/Yobubble/yona-bot/pkg/voicevox_talk"
	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

	"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello this is Yona bot! 🩵",
			},
		})
	},
	"voicevox_talk": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		vvt := &voicevoxTalk.VoicevoxTalk{
			S: s,
			I: i,
			VVTUseCase: &voicevoxTalk.VVTUseCases{
				Vc: s.VoiceConnections[i.GuildID],
			},
		}

		switch options[0].Name {
		case "join":
			vvt.JoinVoiceChannel()
		case "listen":
			start := time.Now()
			vvt.ListenToTheVoiceChannel()
			end := time.Now()
			log.Sugar.Debugf("Take %s seconds", end.Sub(start))
		case "mrs_green_apple_lilac":
			vvt.AudioTest()
		case "disconnect":
			vvt.DisconnectFromTheVoiceChannel()
		}
	},
}
