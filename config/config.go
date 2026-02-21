package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host  string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadDB()  {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err !=nil {
		log.Println("no .env file found, using env variables")
	}
}

func PostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)
}
