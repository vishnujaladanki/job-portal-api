package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIPort      int    `mapstructure:"API_PORT"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout  int    `mapstructure:"IDLE_TIMEOUT"`
	PrivateKey   string `mapstructure:"PRIVATE_KEY"`
	PublicKey    string `mapstructure:"PUBLIC_KEY"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}
