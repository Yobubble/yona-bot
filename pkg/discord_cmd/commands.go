package discordcmd

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "hello",
		Description: "Greeting message.",
	},
	{
		Name:         "voicevox_talk",
		Description:  "A talking bot powered by VOICEVOX Engine",
		DMPermission: new(bool),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "join",
				Description: "Join the voice channel you are currently in.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "listen",
				Description: "Listen for the questions from users in the voice channel.",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "disconnect",
				Description: "Disconnect bot from the voice channel.",
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
