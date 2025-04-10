package vvt

import (
	"fmt"

	"github.com/Yobubble/yona-bot/api"
	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/bwmarrin/discordgo"
)

type VoicevoxTalk struct {
	S          *discordgo.Session
	I          *discordgo.InteractionCreate
	VVTUseCase *VVTUseCases
}

func (v *VoicevoxTalk) AudioTest() {
	if !helper.IsInVoiceChannel(v.S, v.I) {
		log.Sugar.Error("Bot is not in the voice channel")
		return
	}

	v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Playing Lilac by Mrs. GREEN APPLE 🎵",
		},
	})

	// convert from mp3 to dca
	helper.ConvertMp3ToDCA("./assets/audios/vve/Mrs._GREEN_APPLE_Lilac.wav", "./assets/audios/dca/Mrs._GREEN_APPLE_Lilac.dca")

	// play target dca file
	if err := helper.PlayAudio(v.S.VoiceConnections[v.I.GuildID], "./assets/audios/dca/Mrs._GREEN_APPLE_Lilac.dca"); err != nil {
		log.Sugar.Errorf("Error playing audio: %v", err)
		return
	}

}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *VoicevoxTalk) JoinVoiceChannel() {
	log.Sugar.Infof("Attempting voice channel join for user id:%s on channel id:%s", v.I.Member.User.ID, v.I.GuildID)

	guild, err := v.S.Guild(v.I.GuildID)
	if err != nil {
		errMessage := fmt.Sprintf("Cannot get guild info: %v", err)
		log.Sugar.Error(errMessage)
		return
	}

	if err := v.VVTUseCase.createNewChatHistory(guild.Name); err != nil {
		errMessage := fmt.Sprintf("Cannot create chat history: %v", err)
		log.Sugar.Error(errMessage)
		return
	}

	channelId, err := helper.GetUserVoiceChannel(v.S, v.I)
	if err != nil {
		errMessage := fmt.Sprintf("Cannot get user's current channel id: %v", err)
		log.Sugar.Error(errMessage)
		return
	}

	_, err = v.S.ChannelVoiceJoin(v.I.GuildID, channelId, false, false)
	if err != nil {
		errMessage := fmt.Sprintf("Cannot join voice channel: %v", err)
		log.Sugar.Error(errMessage)
		return
	}

	v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Successfully joined the voice channel 💫",
		},
	})
}

func (v *VoicevoxTalk) ListenToTheVoiceChannel() {
	vc := v.S.VoiceConnections[v.I.GuildID]

	if !helper.IsInVoiceChannel(v.S, v.I) {
		log.Sugar.Error("Bot still not in the voice channel")
		return
	}

	v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ready! 💬",
		},
	})

	guild, err := v.S.Guild(v.I.GuildID)
	if err != nil {
		log.Sugar.Errorf("Cannot get guild info: %v", err)
		return
	}

	ssrcs, err := v.VVTUseCase.voiceRecording()
	if err != nil {
		log.Sugar.Errorf("Error recording voice: %v", err)
		return
	}

	oa := api.NewOpenAIAPI(config.C, guild.Name)

	for _, ssrc := range ssrcs {
		ssrcInt := int(ssrc)
		inputPath := fmt.Sprintf("./assets/audios/ssrcs/ogg/%d.ogg", ssrcInt)
		outputPath := fmt.Sprintf("./assets/audios/ssrcs/mp3/%d.mp3", ssrcInt)

		err = helper.ConvertOggToMp3(inputPath, outputPath)
		if err != nil {
			log.Sugar.Errorf("Error converting ogg to mp3:%v", err)
			return
		}

		// STT
		log.Sugar.Debug("Speech To Text...")
		question, err := oa.AudioToText(fmt.Sprintf("./assets/audios/ssrcs/mp3/%d.mp3", ssrc), "これは日本語の音声です。日本語で書き起こしてください。")
		if err != nil {
			log.Sugar.Errorf("Error converting audio to text:%v", err)
			return
		}

		// LLM
		log.Sugar.Debug("LLM Answering...")
		answer, err := oa.AskQuestion(question, "答える前に必ず質問を繰り返す")
		if err != nil {
			log.Sugar.Errorf("Error generating answer: %v", err)
			return
		}

		// TTS
		log.Sugar.Debug("Text To Speech...")
		outputPath = fmt.Sprintf("./assets/audios/vve/mp3/%d.mp3", ssrc)
		err = api.Vve.TextToAudio(answer, outputPath)
		if err != nil {
			log.Sugar.Errorf("Error converting from text to audio: %v", err)
			return
		}

		// mp3 -> dca
		inputPath = fmt.Sprintf("./assets/audios/vve/mp3/%d.mp3", ssrc)
		outputPath = fmt.Sprintf("./assets/audios/dca/%d.dca", ssrc)
		helper.ConvertMp3ToDCA(inputPath, outputPath)

		// play
		log.Sugar.Debug("Playing audio...")
		err = helper.PlayAudio(vc, outputPath)
		if err != nil {
			log.Sugar.Errorf("Error playing audio: %v", err)
			return
		}
	}
}

func (v *VoicevoxTalk) DisconnectFromTheVoiceChannel() {
	if !helper.IsInVoiceChannel(v.S, v.I) {
		v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please use join command first",
			},
		})
	}

	if err := v.S.VoiceConnections[v.I.GuildID].Disconnect(); err != nil {
		log.Sugar.Info("failed to disconnect from voice channel:", err)
		return
	}

	v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Successfully disconnected from the voice channel 💫",
		},
	})
}
