package admin

import (
	"strings"

	"discord-bot/src/controllers"
	"github.com/bwmarrin/discordgo"
)

func extractAmountFromArray(args []string) []string {
	filteredAmount := []string{}

	for _, element := range args {
		if !strings.HasPrefix(element, "<") && len(element) <= 7 {
			filteredAmount = append(filteredAmount, element)
		}
	}

	return filteredAmount
}

func ExecuteGiveSalary(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	salaries := extractAmountFromArray(args)

	for index, user := range m.Mentions {
		controller := controllers.CoinsController{Recipient: user}

		amount := salaries[len(salaries)-1]

		if (len(salaries) - 1) > index {
			amount = salaries[index]
		}

		go controller.UpdateCoins(amount)
	}
}
