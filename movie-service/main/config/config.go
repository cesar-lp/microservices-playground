package config

import (
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Server ServerConfig
	DB     DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Environment string
	Address     string
}

func Load() AppConfig {
	viper.AddConfigPath("../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return buildConfig()
}

func buildConfig() AppConfig {
	return AppConfig{
		Server: ServerConfig{
			Environment: viper.GetString("SERVER_ENVIRONMENT"),
			Address:     viper.GetString("SERVER_ADDRESS"),
		},
		DB: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
		},
	}
}
