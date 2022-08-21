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
}
