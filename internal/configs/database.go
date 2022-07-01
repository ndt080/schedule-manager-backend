package configs

import (
	"fmt"
)

type DatabaseConfig struct {
	Host         string `envconfig:"DB_HOST"`
	Port         string `envconfig:"DB_PORT"`
	Username     string `envconfig:"DB_USERNAME"`
	Password     string `envconfig:"DB_PASSWORD"`
	DatabaseName string `envconfig:"DB_NAME"`
	SSLMode      string `envconfig:"SSL_MODE"`
}

func (config *DatabaseConfig) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DatabaseName, config.Password, config.SSLMode)
}
