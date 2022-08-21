package admin

import (
	"discord-bot/src/config"
	"discord-bot/src/types"
	"github.com/bwmarrin/discordgo"
)

var AdminCommands = map[string]types.Command{
	"ebd": {
		Execute:            ExecuteEmbedSend,
		Description:        "Send embed to channel",
		Params:             []string{"...embed"},
		AllowedPermissions: []int{discordgo.PermissionAdministrator},
		AllowedRoles:       []string{config.Roles.Moderator},
	},
	"ebdedit": {
		Execute:            ExecuteEmbedEdit,
		Description:        "Edit embed",
		Params:             []string{"messageId", "...embed"},
		AllowedPermissions: []int{discordgo.PermissionAdministrator},
		AllowedRoles:       []string{config.Roles.Moderator},
	},
	"salary": {
		Execute:            ExecuteGiveSalary,
		Description:        "Give salary to mentioned users",
		Params:             []string{"...mentions", "...coins"},
		AllowedPermissions: []int{discordgo.PermissionAdministrator},
		AllowedRoles:       []string{config.Roles.Moderator},
	},
}
