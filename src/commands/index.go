package commands

import (
	"discord-bot/src/commands/admin"
	"discord-bot/src/types"
	"golang.org/x/exp/maps"
)

var Commands = map[string]types.Command{}

func InitCommands() {
	maps.Copy(Commands, admin.AdminCommands)
}
