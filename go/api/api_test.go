package api

import (
	"fmt"
	"testing"

	"github.com/Yobubble/yona-bot/config"
)

// func TestAudioToText(t *testing.T) {
// 	c := config.InitConfig()

// 	filePath := "../audios/ja_test/ja_with_little_noise.mp3"

// 	api := NewOpenAIAPIWithKey(c.GetOpenAIAPIKey())

// 	transcription, err := api.AudioToText(filePath)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	fmt.Printf("%s (%T)\n", transcription, transcription)
// }

// func TestNewConversation(t *testing.T) {
// 	c := config.InitConfig()
// 	api := NewOpenAIAPIWithKey(c.GetOpenAIAPIKey())

// 	question := "こんにちは"

// 	answer, err := api.NewConversation(question)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	fmt.Printf("Question: %s", question)
// 	fmt.Printf("%s (%T)\n", answer, answer)
// }

func TestAskNextQuestion(t *testing.T) {
	c := config.InitConfig()

	// Initiate conversation
	api := NewOpenAIAPIWithKey(c.GetOpenAIAPIKey())
	question := "青いツインテールのバーチャルアイドルは誰？" // Who is the virtual idol of blue twin-tail hair?
	answer, err := api.NewConversation(question)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Printf("Question: %s\n", question)
	fmt.Printf("Answer: %s (%T)\n", answer, answer)

	// Ask next question
	question = "彼女が生まれたとき" // When she was borned
	answer, err = api.AskNextQuestion(question)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = api.CloseConversation()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Printf("Question: %s\n", question)
	fmt.Printf("%s (%T)\n", answer, answer)
}
