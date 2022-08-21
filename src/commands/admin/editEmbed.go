package admin

import (
	"encoding/json"
	"strings"

	"discord-bot/src/utils"

	"github.com/bwmarrin/discordgo"
)

func ExecuteEmbedEdit(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := discordgo.MessageEmbed{}

	err := json.Unmarshal([]byte(strings.Join(args[1:], " ")), &embed)

	if err != nil {
		utils.ErrorHandler.Handle(err, "EditEmbed")
		utils.Logger.NotifyUser("Invalid Form Body", m.Author.ID)
		return
	}

	_, err = s.ChannelMessageEditEmbed(m.ChannelID, args[0], &embed)

	if err != nil {
		utils.ErrorHandler.Handle(err, "EditEmbed")
		utils.Logger.NotifyUser("Invalid Form Body", m.Author.ID)
		return
	}
}
