package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is now running")
}
