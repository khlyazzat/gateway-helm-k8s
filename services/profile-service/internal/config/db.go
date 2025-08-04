package config

import (
	"fmt"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	SSLMode    string
}

func (c DBConfig) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBName, c.SSLMode)
}
