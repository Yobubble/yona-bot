package config

import (
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
	lm              enum.LLM
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

func parsedLLM(llm string) enum.LLM {
	switch llm {
	case "GPT4o":
		return enum.GPT4o
	default:
		log.Sugar.Panic("Invalid LLM Model")
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

func (c *Cfg) GetLM() enum.LLM {
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
		log.Sugar.Panicf("Error loading .env file: %v\n", err)
	}

	storage := parsedStorage(os.Getenv("STORAGE"))
	switch storage {
	case enum.Local:
		return &Cfg{
			discordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
			storage:         parsedStorage(os.Getenv("STORAGE")),
			lang:            parsedLang(os.Getenv("LANGUAGE")),
			lm:              parsedLLM(os.Getenv("LM")),
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
			lm:           parsedLLM(os.Getenv("LM")),
			openAIApiKey: os.Getenv("OPENAI_API_KEY"),
			vveBaseUrl:   os.Getenv("VOICEVOX_ENGINE_BASE_URL"),
		}
	default:
		log.Sugar.Panic("Invalid Config")
		return nil
	}
}
