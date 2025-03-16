package main

import (
	"github.com/Yobubble/yona-bot/config"
)

func removeGlobalCommands() {
	config.Sugar.Info("Removing Global commands...")
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		config.Sugar.Panicf("Could not fetch registered commands: %v", err)
	}

	for _, v := range cmds {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			config.Sugar.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
