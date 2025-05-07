package main

import (
	"os"
	"os/signal"

	"github.com/Yobubble/yona-bot/config"
	"github.com/Yobubble/yona-bot/internal/helper"
	"github.com/Yobubble/yona-bot/internal/log"
	discordcmd "github.com/Yobubble/yona-bot/pkg/discord_cmd"
	"github.com/Yobubble/yona-bot/pkg/lm"
	"github.com/Yobubble/yona-bot/pkg/storage"
	"github.com/Yobubble/yona-bot/pkg/stt"
	"github.com/Yobubble/yona-bot/pkg/tts"
	"github.com/bwmarrin/discordgo"
)

var (
	s   *discordgo.Session
	err error
)

func init() {
	log.InitLogger()
	cfg := config.LoadConfig()

	st := storage.SelectStorage(cfg)
	tts := tts.SelectTTSModel(cfg)
	stt := stt.NewOpenAI(cfg)

	ah := helper.AudioHelper{}
	lh := helper.LangHelper{}

	lm := lm.SelectLM(cfg, st, &lh)

	deps := &discordcmd.DepsHolder{
		ST:  st,
		LM:  lm,
		TTS: tts,
		STT: stt,
		AH:  &ah,
	}

	s, err = discordgo.New("Bot " + cfg.GetDiscordBotToken())
	if err != nil {
		log.Sugar.Fatalf("Invalid Bot Parameter: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		dh := helper.NewDiscordHelper(s, i)
		if h, ok := discordcmd.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(dh, deps)
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

	// Create Global Commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(discordcmd.Commands))
	for i, v := range discordcmd.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Sugar.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()
	defer log.Sugar.Sync()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Sugar.Info("Press Ctrl+C to exit")
	<-stop

	log.Sugar.Info("Gracefully shutting down.")
}
