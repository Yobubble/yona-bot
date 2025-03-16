package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Yobubble/yona-bot/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIApi struct {
	C           *openai.Client
	Param       openai.ChatCompletionNewParams
	ChatHistory strings.Builder
}

func (o *OpenAIApi) getConversationTopic() string {
	lines := strings.Split(o.ChatHistory.String(), "\n")
	topic := strings.Split(lines[0], ":")[1]
	return topic
}

func (o *OpenAIApi) updateChatHistory(question string, answer string) {
	questionInfo := fmt.Sprintf("Question:%s\n", question)
	o.ChatHistory.WriteString(questionInfo)
	answerInfo := fmt.Sprintf("Answer:%s\n", answer)
	o.ChatHistory.WriteString(answerInfo)
}

func (o *OpenAIApi) AudioToText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		config.Sugar.Warn("Error opening file")
		return "", err
	}
	defer file.Close()

	transcription, err := o.C.Audio.Transcriptions.New(context.Background(), openai.AudioTranscriptionNewParams{
		File:        openai.F[io.Reader](file),
		Model:       openai.F(openai.AudioModelWhisper1),
		Language:    openai.F("ja"),
		Prompt:      openai.F("これは日本語の音声です。日本語で書き起こしてください。"), // English translation: This is Japanese audio. Please transcribe it in Japanese.
		Temperature: openai.F(0.0),
	})
	if err != nil {
		config.Sugar.Warn("Error transcribing audio")
		return "", err
	}

	return transcription.Text, nil
}

func (o *OpenAIApi) NewConversation(question string) (string, error) {
	o.Param = openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	}

	completion, err := o.C.Chat.Completions.New(context.Background(), o.Param)
	if err != nil {
		config.Sugar.Warn("Error generating answer")
		return "", err
	}

	o.Param.Messages.Value = append(o.Param.Messages.Value, completion.Choices[0].Message)
	o.updateChatHistory(question, completion.Choices[0].Message.Content)

	return completion.Choices[0].Message.Content, nil
}

func (o *OpenAIApi) AskNextQuestion(question string) (string, error) {
	o.Param.Messages.Value = append(o.Param.Messages.Value, openai.UserMessage(question))
	completion, err := o.C.Chat.Completions.New(context.Background(), o.Param)
	if err != nil {
		config.Sugar.Warn("Error generating answer")
		return "", err
	}

	o.updateChatHistory(question, completion.Choices[0].Message.Content)

	return completion.Choices[0].Message.Content, nil
}

func (o *OpenAIApi) CloseConversation() error {
	cvsTopic := o.getConversationTopic()

	outputPath := fmt.Sprintf("./chat_history/%s.txt", cvsTopic)
	err := os.WriteFile(outputPath, []byte(o.ChatHistory.String()), 0644)
	if err != nil {
		config.Sugar.Warn("Error getting conversation topic")
		return err
	}

	return nil
}

func NewOpenAIAPIWithKey(apiKey string) *OpenAIApi {
	return &OpenAIApi{
		C: openai.NewClient(option.WithAPIKey(apiKey)),
	}
}
