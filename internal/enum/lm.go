package enum

import (
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/openai/openai-go"
)

type LM string

const (
	GPT4o LM = "gpt4o"
)

func (l LM) GetOpenAIModel() openai.ChatModel {
	switch l {
	case GPT4o:
		return openai.ChatModelGPT4o
	default:
		log.Sugar.Panic("Invalid OpenAI Model")
		return ""
	}
}
