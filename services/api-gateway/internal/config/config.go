package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPConfig  HTTPConfig
	RedisConfig RedisConfig
	JwtConfig   JwtConfig
}

func New() (*Config, error) {
	accessTTL, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TTL"))
	if err != nil {
		return nil, err
	}
	refreshTTL, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL"))
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		HTTPConfig: HTTPConfig{
			Port: os.Getenv("HTTP_PORT"),
		},
		JwtConfig: JwtConfig{
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
			AccessTTL:     accessTTL,
			RefreshTTL:    refreshTTL,
		},
	}
	return cfg, nil
}
