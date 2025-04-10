package config

import (
	"github.com/Yobubble/yona-bot/internal/log"
	"github.com/spf13/viper"
)

var C *Cfg

type Cfg struct {
	d discord
	o openai
	v voicevoxEngine
}

type discord struct {
	botToken string
}

type openai struct {
	apiKey string
}

type voicevoxEngine struct {
	baseUrl string
}

func (c *Cfg) GetDiscordBotToken() string {
	return c.d.botToken
}

func (c *Cfg) GetOpenAIAPIKey() string {
	return c.o.apiKey
}

func (c *Cfg) GetVoicevoxEngineBaseUrl() string {
	return c.v.baseUrl
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Sugar.Panicf("Error reading config file: %v\n", err)
	}

	C = &Cfg{
		d: discord{
			botToken: viper.GetString("discord.bot_token"),
		},
		o: openai{
			apiKey: viper.GetString("openai.api_key"),
		},
		v: voicevoxEngine{
			baseUrl: viper.GetString("voicevox_engine.base_url"),
		},
	}
}
