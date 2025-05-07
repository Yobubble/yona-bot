package lm

import (
	"context"
	"fmt"
	"strings"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/pkg/storage"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type openAI struct {
	cfg                 *config.Cfg
	conversationHistory []openai.ChatCompletionMessageParamUnion
	st                  storage.Storage
	model               enum.LM
}

func (g *openAI) NewChatHistory(guildName string) error {
	if err := g.st.Write([]byte(""), enum.ChatHistory.GetFullPath(guildName)); err != nil {
		return fmt.Errorf("language model: error wrtting chat history: %w", err)
	}

	return nil
}

func (g *openAI) UpdateChatHistory(guildName string, question string, answer string) error {
	guildPath := enum.ChatHistory.GetFullPath(guildName)

	fileBytes, err := g.st.Read(guildPath)
	if err != nil {
		return fmt.Errorf("language model: error reading chat history: %w", err)
	}

	updatedChatHistory := string(fileBytes) + question + "\n" + answer + "\n"

	err = g.st.Write([]byte(updatedChatHistory), guildPath)
	if err != nil {
		return fmt.Errorf("language model: error writing chat history: %w", err)
	}

	return nil
}

func (g *openAI) AskQuestion(guildName string, question string) (string, error) {
	cli := openai.NewClient(option.WithAPIKey(g.cfg.GetOpenAIAPIKey()))

	err := g.LoadChatHistory(guildName)
	if err != nil {
		return "", fmt.Errorf("language model: error loading chat history: %w", err)
	}

	param := openai.ChatCompletionNewParams{
		Messages: append(g.conversationHistory, openai.UserMessage(question)),
		Seed:     openai.Int(1),
		Model:    g.model.GetOpenAIModel(),
	}

	completion, err := cli.Chat.Completions.New(context.Background(), param)
	if err != nil {
		return "", fmt.Errorf("language model: error getting chat completion: %w", err)
	}

	return completion.Choices[0].Message.Content, nil
}

func (g *openAI) LoadChatHistory(guildName string) error {
	fileBytes, err := g.st.Read(enum.ChatHistory.GetFullPath(guildName))
	if err != nil {
		return fmt.Errorf("language model: error reading chat history: %w", err)
	}

	ch := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	for i, chat := range ch {
		if i%2 == 0 {
			g.conversationHistory = append(g.conversationHistory,
				openai.UserMessage(chat),
			)
		} else {
			g.conversationHistory = append(g.conversationHistory,
				openai.AssistantMessage(chat),
			)
		}
	}

	return nil
}

func newOpenAI(cfg *config.Cfg, st storage.Storage, model enum.LM) LM {
	return &openAI{
		cfg:   cfg,
		st:    st,
		model: model,
	}
}
