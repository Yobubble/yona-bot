package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIApi struct {
	c           *openai.Client
	chatHistory []openai.ChatCompletionMessageParamUnion
	guildName   string
}

func (o *OpenAIApi) updateChatHistory(question string, answer string) error {
	fileBytes, err := os.ReadFile(fmt.Sprintf("./assets/chat_history/%s.txt", o.guildName))
	if err != nil {
		return err
	}

	updatedChatHistory := string(fileBytes) + question + "\n" + answer + "\n"

	err = os.WriteFile(fmt.Sprintf("./assets/chat_history/%s.txt", o.guildName), []byte(updatedChatHistory), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (o *OpenAIApi) AudioToText(filePath string, sysPrompt string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Sugar.Error("Error opening file")
		return "", err
	}
	defer file.Close()

	transcription, err := o.c.Audio.Transcriptions.New(context.Background(), openai.AudioTranscriptionNewParams{
		File:        openai.F[io.Reader](file),
		Model:       openai.F(openai.AudioModelWhisper1),
		Language:    openai.F("ja"),
		Prompt:      openai.F(sysPrompt),
		Temperature: openai.F(0.0),
	})
	if err != nil {
		log.Sugar.Error("Error transcribing audio")
		return "", err
	}

	return transcription.Text, nil
}

func (o *OpenAIApi) AskQuestion(question string, sysPrompt string) (string, error) {
	err := o.loadChatHistory()
	if err != nil {
		log.Sugar.Error("Error load chat history")
		return "", err
	}
	param := openai.ChatCompletionNewParams{
		Messages: openai.F(append(o.chatHistory, openai.SystemMessage(sysPrompt), openai.UserMessage(question))),
		Seed:     openai.Int(1),
		Model:    openai.F(openai.ChatModelGPT4o),
	}

	completion, err := o.c.Chat.Completions.New(context.Background(), param)
	if err != nil {
		log.Sugar.Error("Error generating answer")
		return "", err
	}

	o.updateChatHistory(question, completion.Choices[0].Message.Content)

	return completion.Choices[0].Message.Content, nil
}

func (o *OpenAIApi) loadChatHistory() error {
	fileBytes, err := os.ReadFile(fmt.Sprintf("./assets/chat_history/%s.txt", o.guildName))
	if err != nil {
		return err
	}

	ch := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	for i, chat := range ch {
		if i%2 == 0 {
			o.chatHistory = append(o.chatHistory,
				openai.UserMessage(chat),
			)
		} else {
			o.chatHistory = append(o.chatHistory,
				openai.AssistantMessage(chat),
			)
		}
	}

	return nil
}

func NewOpenAIAPI(cfg *config.Cfg, guildName string) *OpenAIApi {
	apiKey := cfg.GetOpenAIAPIKey()
	return &OpenAIApi{
		c:         openai.NewClient(option.WithAPIKey(apiKey)),
		guildName: guildName,
	}
}
