package tts

import (
	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/log"
)

type TTSModel interface {
	TextToSpeech(text string, outputPath string) error
}

func SelectTTSModel(cfg *config.Cfg) TTSModel {
	switch cfg.GetLang() {
	case enum.JP:
		return newVVE(cfg)
	default:
		log.Sugar.Panic("Select TTS Model Error")
		return nil
	}
}
