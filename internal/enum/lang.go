package enum

import "github.com/Yobubble/yona-bot/internal/log"

type Lang string

const (
	JP Lang = "japanese"
)

func (l Lang) OpenAI() string {
	switch l {
	case JP:
		return "ja"
	default:
		log.Sugar.Panic("Invalid Language")
		return ""
	}
}
