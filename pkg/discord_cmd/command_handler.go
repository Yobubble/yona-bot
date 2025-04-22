package discordcmd

import (
	"github.com/Yobubble/yona-bot/internal/helper"
	vvt "github.com/Yobubble/yona-bot/pkg/voicevox_talk"
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
	"voicevox_talk": func(dh *helper.DiscordHelper, deps *DepsHolder) {
		options := dh.I.ApplicationCommandData().Options

		vvtu := vvt.NewVVTUseCase(deps.ST, deps.LM, deps.TTS, deps.STT, deps.AH)
		vvtc := vvt.NewVVTController(vvtu, dh)

		switch options[0].Name {
		case "join":
			vvtc.JoinVoiceChannel()
		case "listen":
			vvtc.ListenToTheVoiceChannel()
		case "audio_test":
			vvtc.AudioTest()
		case "disconnect":
			vvtc.DisconnectFromTheVoiceChannel()
		}
	},
}
