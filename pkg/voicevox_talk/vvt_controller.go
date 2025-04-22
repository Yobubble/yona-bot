package vvt

import (
	"strconv"

	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/bwmarrin/discordgo"
)

type vvtController struct {
	vvtu *vvtUseCase
	dh   *helper.DiscordHelper
}

func (v *vvtController) AudioTest() {
	if !v.dh.IsBotInVoiceChannel() {
		log.Sugar.Error("Bot is not in the voice channel")
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Playing Lilac by Mrs. GREEN APPLE 🎵",
		}})

	if err := v.dh.PlayAudio(v.dh.GetVoiceConnection(), enum.Audio.GetFullPath("Mrs._GREEN_APPLE_Lilac")); err != nil {
		log.Sugar.Errorf("Error playing audio: %v", err)
		return
	}
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *vvtController) JoinVoiceChannel() {
	guildName, err := v.dh.GetGuildName()
	if err != nil {
		log.Sugar.Errorf("Cannot get guild's name: %v", err)
		return
	}

	if err := v.vvtu.lm.NewChatHistory(guildName); err != nil {
		log.Sugar.Errorf("Cannot create chat history: %v", err)
		return
	}

	channelId, err := v.dh.GetUserVoiceChannel()
	if err != nil {
		log.Sugar.Errorf("Cannot get user's current channel id: %v", err)
		return
	}

	_, err = v.dh.S.ChannelVoiceJoin(v.dh.I.GuildID, channelId, false, false)
	if err != nil {
		log.Sugar.Errorf("Cannot join voice channel: %v", err)
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Joined 💫",
		},
	})
}

func (v *vvtController) ListenToTheVoiceChannel() {
	vc := v.dh.GetVoiceConnection()
	guildName, err := v.dh.GetGuildName()
	if err != nil {
		log.Sugar.Errorf("Cannot get guild's name: %v", err)
		return
	}

	if !v.dh.IsBotInVoiceChannel() {
		log.Sugar.Error("Bot is not in the voice channel")
		return
	}

	if err := v.vvtu.pre(); err != nil {
		log.Sugar.Errorf("Preparing Bot Failed: %v", err)
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ready! 💬",
		},
	})

	ssrcs, err := v.vvtu.voiceRecording(vc)
	if err != nil {
		log.Sugar.Errorf("Error recording voice: %v", err)
		return
	}

	for _, ssrc := range ssrcs {
		ssrcStr := strconv.Itoa(int(ssrc))

		if err := v.vvtu.processRecordAudio(ssrcStr, guildName); err != nil {
			log.Sugar.Errorf("Processing Recording Audios Error: %v", err)
		}

		if err := v.dh.PlayAudio(vc, enum.Audio.GetFullPath(ssrcStr)); err != nil {
			log.Sugar.Errorf("Error playing audio: %v", err)
			return
		}
	}
}

func (v *vvtController) DisconnectFromTheVoiceChannel() {
	if !v.dh.IsBotInVoiceChannel() {
		v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please use join command first",
			},
		})
		return
	}

	if err := v.vvtu.post(); err != nil {
		log.Sugar.Errorf("Cleaning Failed: %v", err)
		return
	}

	if err := v.dh.GetVoiceConnection().Disconnect(); err != nil {
		log.Sugar.Errorf("failed to disconnect from voice channel: %v", err)
		return
	}

	v.dh.S.InteractionRespond(v.dh.I.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Disconnected 💫",
		},
	})
}

func NewVVTController(vvtu *vvtUseCase, dh *helper.DiscordHelper) *vvtController {
	return &vvtController{
		vvtu: vvtu,
		dh:   dh,
	}
}
