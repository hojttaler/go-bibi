package handlers

import (
	"strings"
	"time"

	"discord-bot/src/commands"
	"discord-bot/src/config"
	"discord-bot/src/types"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slices"
)

func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Content) == 0 {
		return
	}

	if m.Content[0:1] != config.Config.Bot.Prefix {
		return
	}

	go func() {
		time.Sleep(2 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	}()

	commandArr := strings.Split(m.Content[1:], " ")

	commandName := commandArr[0]
	args := commandArr[1:]

	command := commands.Commands[commandName]

	if command.Execute == nil {
		return
	}

	if !isPassedTests(s, m, command) {
		return
	}

	go command.Execute(s, m, args)
}

func isPassedTests(s *discordgo.Session, m *discordgo.MessageCreate, command types.Command) bool {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	if err != nil {
		return false
	}

	return isPassedPermissionsAndRolesTest(m.Member, int(p), command) && isPassedChannelsTest(m.ChannelID, command)
}

func isPassedPermissionsAndRolesTest(member *discordgo.Member, memberPermissions int, command types.Command) bool {
	if (command.AllowedPermissions == nil && command.AllowedRoles == nil) ||
		(len(command.AllowedPermissions) == 0 && len(command.AllowedRoles) == 0) {
		return true
	}

	isPassed := false

	if command.AllowedPermissions != nil && len(command.AllowedPermissions) != 0 {
		for _, permission := range command.AllowedPermissions {
			if memberPermissions&permission == permission {
				isPassed = true
			}
		}
	}

	if command.AllowedRoles != nil && len(command.AllowedRoles) != 0 {
		for _, role := range command.AllowedRoles {
			if slices.Contains(member.Roles, role) {
				isPassed = true
			}
		}
	}

	return isPassed
}

func isPassedChannelsTest(channelId string, command types.Command) bool {
	if command.AllowedChannels == nil || len(command.AllowedChannels) == 0 {
		return true
	}

	return slices.Contains(command.AllowedChannels, channelId)
}
