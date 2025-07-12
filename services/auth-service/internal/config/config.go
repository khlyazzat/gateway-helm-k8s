package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBConfig    DBConfig
	HTTPConfig  HTTPConfig
	RedisConfig RedisConfig
	JwtConfig   JwtConfig
}

func New() (*Config, error) {
	log.Println("os.Getenv", os.Getenv("DB_USER"))
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}
	accessTTL, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TTL"))
	if err != nil {
		return nil, err
	}
	refreshTTL, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL"))
	if err != nil {
		return nil, err
	}
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
		RedisConfig: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
		JwtConfig: JwtConfig{
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
			AccessTTL:     accessTTL,
			RefreshTTL:    refreshTTL,
		},
	}

	return cfg, nil
}
