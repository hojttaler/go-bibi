package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"discord-bot/src/config"
	"discord-bot/src/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/vicanso/go-axios"
)

func sendUpdateCoinsPostRequest(userId string, amount string) error {
	headers := http.Header{}
	headers.Set("authorization", config.Config.EightNine.Password)

	response, err := axios.Request(&axios.Config{
		URL:     config.Config.EightNine.Endpoint,
		Route:   "/givecoins/:userId/:amount",
		Params:  map[string]string{userId: userId, amount: amount},
		Method:  http.MethodPost,
		Headers: headers,
	})

	if err != nil {
		return err
	}

	fmt.Println(response.Status)

	if response.Status != http.StatusOK {
		return errors.New("EightNine receive request error")
	}

	return nil
}

type CoinsController struct {
	Sender    *discordgo.User
	Recipient *discordgo.User
	Client    *discordgo.Session
}

func (c *CoinsController) UpdateCoins(amount string) error {
	return sendUpdateCoinsPostRequest(c.Recipient.ID, amount)
}

func (c *CoinsController) WriteLogs() {
	logContent := fmt.Sprintf("<@%s> выдал %s единиц(у) коинов пользователю <@%s>")

	utils.Logger.Log(logContent, config.Channels.Logs.Salary)
}
