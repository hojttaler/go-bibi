package types

type Roles struct {
	Moderator          string `env:"ROLE_MODERATOR"`
	SeniorRedactor     string `env:"ROLE_SENIOR_REDACTOR"`
	Redactor           string `env:"ROLE_REDACTOR"`
	SeniorEventManager string `env:"ROLE_SENIOR_EVENT_MANAGER"`
	EventManager       string `env:"ROLE_EVENT_MANAGER"`
}
