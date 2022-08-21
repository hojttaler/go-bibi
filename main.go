package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"discord-bot/src/commands"
	"discord-bot/src/config"
	"discord-bot/src/events"
	"discord-bot/src/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Get config
	err := config.LoadConfig("dev")

	if err != nil {
		log.Fatal("Can't load config: ", err)
	}

	// Create bot
	discord, err := discordgo.New("Bot " + config.Config.Bot.Token)

	// Load logger
	utils.Logger = utils.LogsHandler{
		Client: discord,
	}

	// Load error handler
	utils.ErrorHandler = utils.ErrorHandleLogger{
		Logger: utils.Logger,
	}

	if err != nil {
		utils.ErrorHandler.Handle(err, "BotCreate")

		return
	}

	// Add intent to bot
	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Init commands and events
	events.InitEvents(discord)
	commands.InitCommands()

	// Open discord connection
	if err := discord.Open(); err != nil {
		utils.ErrorHandler.Handle(err, "BotConnectionOpen")

		panic(err)
	}

	// Close discord connection on Ctrl-C command
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
