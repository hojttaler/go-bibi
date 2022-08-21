package types

import "github.com/bwmarrin/discordgo"

type ExecuteFunc func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)

type Command struct {
	Execute            ExecuteFunc
	Description        string
	Params             []string
	AllowedRoles       []string
	AllowedPermissions []int
	AllowedChannels    []string
}
