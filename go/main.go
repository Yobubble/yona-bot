package main

import (
	"os"
	"os/signal"

	"github.com/Yobubble/yona-bot/cmd"
	"github.com/Yobubble/yona-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	s   *discordgo.Session
	err error
)

func init() {
	c := config.InitConfig()
	config.InitLogger()

	s, err = discordgo.New("Bot " + c.GetDiscordBotToken())
	if err != nil {
		config.Sugar.Fatalf("Invalid Bot Parameter: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := cmd.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go
func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		config.Sugar.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		config.Sugar.Fatalf("Cannot open the session: %v", err)
	}

	config.Sugar.Info("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(cmd.Commands))

	// Create Global Command
	for i, v := range cmd.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			config.Sugar.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	config.Sugar.Info("Press Ctrl+C to exit")
	<-stop

	// NOTE: remove all the global commands
	// removeGlobalCommands()

	config.Sugar.Info("Gracefully shutting down.")
}
