package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host           string
	Port           int
	TrustedProxies []string
}

func New() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	error := viper.ReadInConfig()
	if error != nil {
		panic(fmt.Errorf("fatal error config file: %w", error))
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode config file to struct: %w", err))
	}

	return &config
}
