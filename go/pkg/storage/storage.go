package storage

import (
	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/log"
)

type Storage interface {
	Write(data []byte, filePath string) error
	Read(filePath string) ([]byte, error)
}

func SelectStorage(cfg *config.Cfg) Storage {
	switch cfg.GetStorage() {
	case enum.Local:
		return newLocal()
	case enum.S3:
		s3 := cfg.GetS3Config()
		return newS3(&s3)
	default:
		log.Sugar.Panic("Select Storage Error")
		return nil
	}
}
