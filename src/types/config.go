package types

type Config struct {
	Bot struct {
		Token  string `env:"BOT_TOKEN"`
		Prefix string `env:"BOT_PREFIX"`
	}

	Database struct {
		Link string `env:"DATABASE_LINK"`
	}

	EightNine struct {
		Endpoint string `env:"EIGHTNINE_ENDPOINT"`
		Password string `env:"EIGHTNINE_PASSWORD"`
	}

	Channels struct {
		Logs struct {
			Errors string `env:"CHANNEL_LOGS_ERRORS"`
			Salary string `env:"CHANNEL_LOGS_SALARY"`
		}
	}

	Roles struct {
		Moderator          string `env:"ROLE_MODERATOR"`
		SeniorRedactor     string `env:"ROLE_SENIOR_REDACTOR"`
		Redactor           string `env:"ROLE_REDACTOR"`
		SeniorEventManager string `env:"ROLE_SENIOR_EVENT_MANAGER"`
		EventManager       string `env:"ROLE_EVENT_MANAGER"`
	}
}
