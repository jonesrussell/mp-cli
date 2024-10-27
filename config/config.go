package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string
	APIURL   string
	Username string
	Password string
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

	return &config, nil
}
