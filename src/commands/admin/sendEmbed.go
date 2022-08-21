package admin

import (
	"encoding/json"
	"strings"

	"discord-bot/src/utils"

	"github.com/bwmarrin/discordgo"
)

func ExecuteEmbedSend(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := discordgo.MessageEmbed{}

	err := json.Unmarshal([]byte(strings.Join(args, " ")), &embed)

	if err != nil {
		utils.ErrorHandler.Handle(err, "SendEmbed")
		utils.Logger.NotifyUser("Invalid Form Body", m.Author.ID)
		return
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embed)

	if err != nil {
		utils.ErrorHandler.Handle(err, "SendEmbed")
		utils.Logger.NotifyUser("Invalid Form Body", m.Author.ID)
		return
	}
}
