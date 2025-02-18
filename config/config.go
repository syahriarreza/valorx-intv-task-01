package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDSN string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("DATABASE_DSN", "postgres://user:password@localhost:5432/dbname?sslmode=disable")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	databaseDSN := viper.GetString("DATABASE_DSN")
	if databaseDSN == "" {
		log.Fatal("DATABASE_DSN configuration is required")
	}

	return &Config{
		DatabaseDSN: databaseDSN,
	}
}
