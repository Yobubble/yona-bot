package config

import (
	"github.com/spf13/viper"
)

type config struct {
	d discord
	o openai
}

type discord struct {
	botToken string
}

type openai struct {
	apiKey string
}

func (c *config) GetDiscordBotToken() string {
	return c.d.botToken
}

func (c *config) GetOpenAIAPIKey() string {
	return c.o.apiKey
}

func InitConfig() *config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		Sugar.Panicf("Error reading config file: %v\n", err)
	}

	return &config{
		d: discord{
			botToken: viper.GetString("discord.bot_token"),
		},
		o: openai{
			apiKey: viper.GetString("openai.api_key"),
		},
	}
}
