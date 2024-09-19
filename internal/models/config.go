package models

type Config struct {
	Next string
	Previous any
}

func NewConfig() (*Config) {
	return &Config{
		Next: "",
		Previous: "",
	}
}