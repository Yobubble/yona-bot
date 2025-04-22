package discordcmd

import (
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/pkg/lm"
	"github.com/Yobubble/yona-bot/pkg/storage"
	"github.com/Yobubble/yona-bot/pkg/stt"
	"github.com/Yobubble/yona-bot/pkg/tts"
)

type DepsHolder struct {
	ST  storage.Storage
	LM  lm.LM
	TTS tts.TTSModel
	STT stt.STTModel
	AH  *helper.AudioHelper
}
