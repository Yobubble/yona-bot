package lm

import (
	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/Yobubble/yona-bot/pkg/storage"
)

type LM interface {
	AskQuestion(guildName string, question string) (string, error)
	NewChatHistory(guildName string) error
	LoadChatHistory(guildName string) error
	UpdateChatHistory(guildName string, question string, answer string) error
}

func SelectLM(cfg *config.Cfg, st storage.Storage, lh *helper.LangHelper) LM {
	switch cfg.GetLM() {
	case enum.GPT4o:
		return newOpenAI(cfg, st, lh, enum.GPT4o)
	default:
		log.Sugar.Panic("Select LM Invalid")
		return nil
	}

}
