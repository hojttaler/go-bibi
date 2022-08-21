package events

import "github.com/bwmarrin/discordgo"

func InitEvents(s *discordgo.Session) {
	s.AddHandler(Ready)
	s.AddHandler(MessageCreate)
}
