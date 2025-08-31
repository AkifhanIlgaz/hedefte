package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MongoConfig struct {
	ConnString string `mapstructure:"connection_string"`
}

type PostgresConfig struct {
	ConnString string `mapstructure:"connection_string"`
}

type AppConfig struct {
	Mongo    MongoConfig    `mapstructure:"mongo"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

func Load() (AppConfig, error) {
	var config AppConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config/")

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil
}
