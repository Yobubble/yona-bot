package cmd

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "hello",
		Description: "Hello! This is Yona!",
	},
	{
		Name:         "voicevox_talk",
		Description:  "A talking bot powered by VOICEVOX Engine",
		DMPermission: new(bool),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "join",
				Description: "Join the voice channel you are in",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "listen",
				Description: "Listen to the voice channel you are in for 10 seconds",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "disconnect",
				Description: "Safe disconnect from the voice channel",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
}
