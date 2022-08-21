package utils

import (
	"fmt"

	"discord-bot/src/config"
	"github.com/bwmarrin/discordgo"
)

var Logger LogsHandler

type LogsHandler struct {
	Client *discordgo.Session
}

func (l *LogsHandler) sendLogToConsole(err error) {
	fmt.Printf("невозможно отправить сообщение: %s\n", err.Error())
}

func (l *LogsHandler) LogError(error error) {
	_, sendError := l.Client.ChannelMessageSend(config.Config.Channels.Logs.Errors, error.Error())

	if sendError != nil {
		l.sendLogToConsole(error)
		l.sendLogToConsole(sendError)
	}
}

func (l *LogsHandler) Log(content string, channel string) {
	_, err := l.Client.ChannelMessageSend(channel, content)

	if err != nil {
		ErrorHandler.Handle(err, "SendLog")
	}
}

func (l *LogsHandler) NotifyUser(content string, userId string) {
	channel, err := l.Client.UserChannelCreate(userId)

	if err != nil {
		ErrorHandler.Handle(err, "NotifyUser")
		return
	}

	_, err = l.Client.ChannelMessageSend(channel.ID, content)

	if err != nil {
		ErrorHandler.Handle(err, "NotifyUser")
		return
	}
}
