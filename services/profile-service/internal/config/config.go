package config

import (
	"os"
)

type Config struct {
	DBConfig   DBConfig
	HTTPConfig HTTPConfig
}

func New() (*Config, error) {
	cfg := &Config{
		DBConfig: DBConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBName:     os.Getenv("DB_NAME"),
			SSLMode:    os.Getenv("DB_SSL_MODE"),
		},
		HTTPConfig: HTTPConfig{
			Port: os.Getenv("HTTP_PORT"),
		},
	}

	return cfg, nil
}
