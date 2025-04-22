package stt

import (
	"github.com/Yobubble/yona-bot/internal/enum"
)

type STTModel interface {
	AudioToText(filePath string, lang enum.Lang) (string, error)
}
