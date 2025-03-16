package voicevoxTalk

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/Yobubble/yona-bot/config"
	"github.com/bwmarrin/discordgo"
)

type VoiceVoxTalk struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/voice_receive/main.go
func (v *VoiceVoxTalk) JoinVoiceChannel() {
	config.Sugar.Infof("Attempting voice channel join for user id:%s on channel id:%s", v.I.Member.User.ID, v.I.GuildID)

	channelId, err := getUserVoiceChannel(v.S, v.I)
	if err != nil {
		config.Sugar.Error("Cannot get user's current channel id:", zap.Error(err))
		return
	}

	_, err = v.S.ChannelVoiceJoin(v.I.GuildID, channelId, false, false)
	if err != nil {
		config.Sugar.Error("Couldn't join a channel:", zap.Error(err))
		return
	}
}

func (v *VoiceVoxTalk) ListenToTheVoiceChannel() {
	if !isInVoiceChannel(v.S, v.I) {
		config.Sugar.Error("Bot still not in the voice channel")
	}

	ssrcs, err := voiceRecording(v.S.VoiceConnections[v.I.GuildID])
	if err != nil {
		config.Sugar.Error("Error recording voice:", zap.Error(err))
	}

	for ssrc := range ssrcs {
		err = convertOggToMp3(strconv.FormatUint(uint64(ssrc), 10))
		if err != nil {
			config.Sugar.Error("Error converting ogg to mp3:", zap.Error(err))
		}
	}

	// NOTE: STT

	// NOTE: LLM
	// NOTE: VVE
}

func (v *VoiceVoxTalk) DisconnectFromTheVoiceChannel() {
	// NOTE: Check if the bot is already in the channel
	if !isInVoiceChannel(v.S, v.I) {
		v.S.InteractionRespond(v.I.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please use join command first",
			},
		})
	}

	err := v.S.VoiceConnections[v.I.GuildID].Disconnect()
	if err != nil {
		config.Sugar.Info("failed to disconnect from voice channel:", err)
		return
	}
}
