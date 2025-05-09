package discordcmd

import (
	"github.com/Yobubble/yona-bot/internal/helper"
	cvs "github.com/Yobubble/yona-bot/pkg/conversation"
	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(dh *helper.DiscordHelper, deps *DepsHolder){

	"hello": func(dh *helper.DiscordHelper, deps *DepsHolder) {
		dh.S.InteractionRespond(dh.I.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello this is Yona bot! 🩵",
			},
		})
	},
	"conversation": func(dh *helper.DiscordHelper, deps *DepsHolder) {
		options := dh.I.ApplicationCommandData().Options

		vvtu := cvs.NewCVSUseCase(deps.ST, deps.LM, deps.TTS, deps.STT, deps.AH)
		vvtc := cvs.NewCVSController(vvtu, dh)

		switch options[0].Name {
		case "join":
			vvtc.JoinVoiceChannel()
		case "listen":
			vvtc.ListenToTheVoiceChannel()
		case "disconnect":
			vvtc.DisconnectFromTheVoiceChannel()
		case "audio_test":
			vvtc.AudioTest()
		}
	},
}
