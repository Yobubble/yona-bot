package stt

import (
	"context"
	"fmt"
	"os"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
)

type openAI struct {
	c    openai.Client
	lang enum.Lang
}

func (w *openAI) AudioToText(filePath string, lang enum.Lang) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("speech to text: failed to open file: %w", err)
	}
	defer file.Close()

	transcription, err := w.c.Audio.Transcriptions.New(context.Background(), openai.AudioTranscriptionNewParams{
		File:  file,
		Model: openai.AudioModelGPT4oTranscribe,
		Language: param.Opt[string]{
			Value: lang.OpenAI(),
		},
	})
	if err != nil {
		return "", fmt.Errorf("speech to text: failed to transcribe audio: %w", err)
	}

	return transcription.Text, nil
}

func NewOpenAI(cfg *config.Cfg) STTModel {
	return &openAI{
		c:    openai.NewClient(option.WithAPIKey(cfg.GetOpenAIAPIKey())),
		lang: cfg.GetLang(),
	}
}
