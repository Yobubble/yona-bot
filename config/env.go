package config

import (
	"errors"
	"os"

	"github.com/Yobubble/yona-bot/internal/enum"
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/joho/godotenv"
)

type Cfg struct {
	discordBotToken string
	storage         enum.Storage
	s3              S3Config
	lang            enum.Lang
	lm              enum.LM
	openAIApiKey    string
	vveBaseUrl      string
}

type S3Config struct {
	S3Bucket           string
	AWSAccessKeyId     string
	AWSSecretAccessKey string
	AWSRegion          string
}

func parsedStorage(storage string) enum.Storage {
	switch storage {
	case "Local":
		return enum.Local
	case "S3":
		return enum.S3
	default:
		log.Sugar.Panic("Invalid Storage")
		return ""
	}
}

func parsedLang(lang string) enum.Lang {
	switch lang {
	case "JP":
		return enum.JP
	default:
		log.Sugar.Panic("Invalid Language")
		return ""
	}
}

func parsedLM(lm string) enum.LM {
	switch lm {
	case "GPT4o":
		return enum.GPT4o
	default:
		log.Sugar.Panic("Invalid LM")
		return ""
	}
}

func (c *Cfg) GetDiscordBotToken() string {
	return c.discordBotToken
}

func (c *Cfg) GetStorage() enum.Storage {
	return c.storage
}

func (c *Cfg) GetS3Config() S3Config {
	return c.s3
}

func (c *Cfg) GetLang() enum.Lang {
	return c.lang
}

func (c *Cfg) GetLM() enum.LM {
	return c.lm
}

func (c *Cfg) GetOpenAIAPIKey() string {
	return c.openAIApiKey
}

func (c *Cfg) GetVVEBaseUrl() string {
	return c.vveBaseUrl
}

func LoadConfig() *Cfg {
	err := godotenv.Load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Sugar.Info("No .env file found, rely on system environment variable.")
		} else {
			log.Sugar.Panicf("Error loading .env file: %v\n", err)
		}
	}

	storage := parsedStorage(os.Getenv("STORAGE"))
	switch storage {
	case enum.Local:
		return &Cfg{
			discordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
			storage:         parsedStorage(os.Getenv("STORAGE")),
			lang:            parsedLang(os.Getenv("LANGUAGE")),
			lm:              parsedLM(os.Getenv("LM")),
			openAIApiKey:    os.Getenv("OPENAI_API_KEY"),
			vveBaseUrl:      os.Getenv("VOICEVOX_ENGINE_BASE_URL"),
		}
	case enum.S3:
		return &Cfg{
			discordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
			storage:         parsedStorage(os.Getenv("STORAGE")),
			s3: S3Config{
				S3Bucket:           os.Getenv("S3_BUCKET"),
				AWSAccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
				AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
				AWSRegion:          os.Getenv("AWS_REGION"),
			},
			lang:         parsedLang(os.Getenv("LANGUAGE")),
			lm:           parsedLM(os.Getenv("LM")),
			openAIApiKey: os.Getenv("OPENAI_API_KEY"),
			vveBaseUrl:   os.Getenv("VOICEVOX_ENGINE_BASE_URL"),
		}
	default:
		log.Sugar.Panic("Invalid Environment Variable")
		return nil
	}
}
