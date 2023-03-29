package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	Driver string
	URL    string
    DNS    string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
