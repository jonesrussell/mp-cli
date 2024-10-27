package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string
	APIBaseURL string
	Username   string
	Password   string
	UserID     string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.mp-cli")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	// Set default values if not provided
	if config.APIBaseURL == "" {
		config.APIBaseURL = "http://localhost:8080"
	}

	// Validate required fields
	if config.Username == "" || config.Password == "" {
		return nil, fmt.Errorf("username and password are required in the config file")
	}

	return &config, nil
}
