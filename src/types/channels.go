package types

type Channels struct {
	Logs struct {
		Errors string `env:"CHANNEL_LOGS_ERRORS"`
		Salary string `env:"CHANNEL_LOGS_SALARY"`
	}
}
