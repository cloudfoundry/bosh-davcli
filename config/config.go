package config

type Config struct {
	User          string
	Password      string
	Endpoint      string
	RetryAttempts uint
	CA            string
}
