package cmd

import (
	voicevoxTalk "github.com/Yobubble/yona-bot/voicevox_talk"
	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello World",
			},
		})
	},
	"voicevox_talk": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		vvt := &voicevoxTalk.VoiceVoxTalk{
			S: s,
			I: i,
		}
		switch options[0].Name {
		case "join":
			vvt.JoinVoiceChannel()
			break
		case "listen":
			vvt.ListenToTheVoiceChannel()
			break
		case "disconnect":
			vvt.DisconnectFromTheVoiceChannel()
			break
		}
	},
}
