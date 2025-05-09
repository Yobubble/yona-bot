package discordcmd

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "hello",
		Description: "Greeting message.",
	},
	{
		Name:         "conversation",
		Description:  "Let's have a conversation!",
		DMPermission: new(bool),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "join",
				Description: "Join the voice channel you are currently in.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "listen",
				Description: "Listen for the questions in the voice channel and stop when there is silence.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "disconnect",
				Description: "Remove recorded files and Disconnect bot from the voice channel.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "audio_test",
				Description: "Play Lilac by Mrs. GREEN APPLE for the audio testing purpose.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
}
