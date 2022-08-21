package config

import (
	"fmt"
	"os"

	"discord-bot/src/types"
	"github.com/golobby/dotenv"
)

var Config types.Config
var Roles types.Roles
var Channels types.Channels

func LoadConfig(env string) error {
	envFile := ".env"

	if env == "prod" {
		envFile = ".env"
	} else {
		envFile = ".env.dev"
	}

	fmt.Println("Loaded", envFile, "file as config")

	file, err := os.Open(envFile)

	if err != nil {
		return err
	}

	err = dotenv.NewDecoder(file).Decode(&Config)
	if err != nil {
		return err
	}

	err = dotenv.NewDecoder(file).Decode(&Roles)
	if err != nil {
		return err
	}

	err = dotenv.NewDecoder(file).Decode(&Channels)
	if err != nil {
		return err
	}

	return nil
}
