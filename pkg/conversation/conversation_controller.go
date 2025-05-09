package cvs

import (
	"strconv"

	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type cvsController struct {
	cvsu *cvsUseCase
	dh   *helper.DiscordHelper
}

func (v *cvsController) AudioTest() {
	if !v.dh.IsBotInVoiceChannel() {
		log.Sugar.Error("Bot is not in the voice channel")
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Playing Lilac by Mrs. GREEN APPLE 🎵",
		}})

	if err := v.dh.PlayAudio(v.dh.GetVoiceConnection(), "./asset/audios/Mrs._GREEN_APPLE_Lilac.dca"); err != nil {
		log.Sugar.Error("Error playing audio: ", zap.Error(err))
		return
	}
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *cvsController) JoinVoiceChannel() {
	guildName, err := v.dh.GetGuildName()
	if err != nil {
		log.Sugar.Error("Cannot get discord server's name: ", zap.Error(err))
		return
	}

	if err := v.cvsu.lm.NewChatHistory(guildName); err != nil {
		log.Sugar.Error("Cannot create new chat history: ", zap.Error(err))
		return
	}

	channelId, err := v.dh.GetUserVoiceChannel()
	if err != nil {
		log.Sugar.Error("Cannot get user's current voice channel: ", zap.Error(err))
		return
	}

	_, err = v.dh.S.ChannelVoiceJoin(v.dh.I.GuildID, channelId, false, false)
	if err != nil {
		log.Sugar.Error("Cannot join voice channel: ", zap.Error(err))
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Joined 💫",
		},
	})
}

func (v *cvsController) ListenToTheVoiceChannel() {
	vc := v.dh.GetVoiceConnection()
	guildName, err := v.dh.GetGuildName()
	if err != nil {
		log.Sugar.Error("Cannot get guild's name: ", zap.Error(err))
		return
	}

	if !v.dh.IsBotInVoiceChannel() {
		log.Sugar.Error("Bot is not in the voice channel")
		return
	}

	if err := v.cvsu.pre(); err != nil {
		log.Sugar.Error("Preparing Bot Failed: ", zap.Error(err))
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ready! 💬",
		},
	})

	ssrcs, err := v.cvsu.voiceRecording(vc)
	if err != nil {
		log.Sugar.Error("Error recording voice: ", zap.Error(err))
		return
	}

	for _, ssrc := range ssrcs {
		ssrcStr := strconv.Itoa(int(ssrc))

		if err := v.cvsu.processRecordAudio(ssrcStr, guildName); err != nil {
			log.Sugar.Error("Processing Recording Audios Error: ", zap.Error(err))
			return
		}

		if err := v.dh.PlayAudio(vc, enum.Audio.GetFullPath(ssrcStr)); err != nil {
			log.Sugar.Error("Error playing audio: ", zap.Error(err))
			return
		}
	}
}

func (v *cvsController) DisconnectFromTheVoiceChannel() {
	if !v.dh.IsBotInVoiceChannel() {
		v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please use join command first",
			},
		})
		return
	}

	if err := v.cvsu.post(); err != nil {
		log.Sugar.Error("Cleaning Failed: ", zap.Error(err))
		return
	}

	if err := v.dh.GetVoiceConnection().Disconnect(); err != nil {
		log.Sugar.Error("failed to disconnect from voice channel: ", zap.Error(err))
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Disconnected 💫",
		},
	})
}

func NewCVSController(cvsu *cvsUseCase, dh *helper.DiscordHelper) *cvsController {
	return &cvsController{
		cvsu: cvsu,
		dh:   dh,
	}
}
