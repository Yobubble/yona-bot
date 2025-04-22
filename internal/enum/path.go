package enum

import (
	"fmt"

	"github.com/Yobubble/yona-bot/internal/log"
)

type Path string

const (
	ChatHistory Path = "./assets/chat_histories/"
	SSRC_OGG    Path = "./assets/audios/ssrcs_ogg/"
	SSRC_MP3    Path = "./assets/audios/ssrcs_mp3/"
	VVE         Path = "./assets/audios/vves/"
	Audio       Path = "./assets/audios/audios/"
)

func (p Path) GetFullPath(fileName string) string {
	const base = "%s%s%s"
	switch p {
	case ChatHistory:
		return fmt.Sprintf(base, p, fileName, ChatHistory.GetFormat())
	case SSRC_OGG:
		return fmt.Sprintf(base, p, fileName, SSRC_OGG.GetFormat())
	case SSRC_MP3:
		return fmt.Sprintf(base, p, fileName, SSRC_MP3.GetFormat())
	case VVE:
		return fmt.Sprintf(base, p, fileName, VVE.GetFormat())
	case Audio:
		return fmt.Sprintf(base, p, fileName, Audio.GetFormat())
	default:
		log.Sugar.Panic("Invalid Path")
		return ""
	}
}

func (p Path) GetPath() string {
	return string(p)
}

func (p Path) GetFormat() string {
	switch p {
	case ChatHistory:
		return ".txt"
	case SSRC_OGG:
		return ".ogg"
	case SSRC_MP3:
		return ".mp3"
	case VVE:
		return ".mp3"
	case Audio:
		return ".dca"
	default:
		log.Sugar.Panic("Invalid Format")
		return ""
	}
}
