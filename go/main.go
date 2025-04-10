package main

import (
	"os"
	"os/signal"

	"github.com/Yobubble/yona-bot/api"
	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/log"
	discordcmd "github.com/Yobubble/yona-bot/pkg/discord_cmd"
	"github.com/bwmarrin/discordgo"
)

var (
	s   *discordgo.Session
	err error
)

func init() {
	config.InitConfig()
	log.InitLogger()
	api.InitVoicevoxEngineApi(config.C)

	s, err = discordgo.New("Bot " + config.C.GetDiscordBotToken())
	if err != nil {
		log.Sugar.Fatalf("Invalid Bot Parameter: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := discordcmd.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

// Ref: https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go
func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Sugar.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the session: %v", err)
	}

	log.Sugar.Info("Adding commands...")

	// Create Global Command
	registeredCommands := make([]*discordgo.ApplicationCommand, len(discordcmd.Commands))
	for i, v := range discordcmd.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Sugar.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Sugar.Info("Press Ctrl+C to exit")
	<-stop

	// NOTE: remove all the global commands - TODO: change to script
	// removeGlobalCommands()

	err := api.Vve.CloseVoicevox()
	if err != nil {
		log.Sugar.Errorf("Cannot close Voicevox: %v", err)
	}
	log.Sugar.Info("Gracefully shutting down.")
}
