package config

type Config struct {
	User          string
	Password      string
	Endpoint      string
	RetryAttempts uint
	CaCert        string `json:"ca_cert"`
}
