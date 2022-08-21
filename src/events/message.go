package events

import (
	"discord-bot/src/handlers"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	handlers.CommandHandler(s, m)
}
