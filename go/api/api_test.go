package api

import (
	"fmt"
	"testing"

	"github.com/Yobubble/yona-bot/config"
)

func TestOpenAIApi(t *testing.T) {
	config.InitConfig()
	oa := NewOpenAIAPI(config.C, "test1")

	fmt.Println("Test Audio To Text...")
	res, err := oa.AudioToText("../../audios/ja_test/ja_with_little_noise.mp3", "これは日本語の音声です。日本語で書き起こしてください。")
	if err != nil {
		t.Errorf("Error Converting Audio To Text: %v", err)
	}
	fmt.Println("Text:", res)

	fmt.Println("Test Asking First Question...")
	answer, err := oa.AskQuestion("青いツインテールのバーチャルアイドルは誰？", "答える前に必ず質問を繰り返す")
	if err != nil {
		t.Errorf("Error Asking Question: %v", err)
	}
	fmt.Println(answer)

	fmt.Println("Test Asking Second Question...")
	answer, err = oa.AskQuestion("彼女が生まれたとき", "答える前に必ず質問を繰り返す")
	if err != nil {
		t.Errorf("Error Asking Question: %v", err)
	}
	fmt.Println(answer)
}
