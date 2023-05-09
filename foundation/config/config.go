package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPClient HTTPClient
	HTTPServer HTTPServer
	SQLite     SQLite
}

type HTTPClient struct {
	Timeout time.Duration
}

type HTTPServer struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type SQLite struct {
	Path string
}

// Load loads the configuration from the config file.
func Load() (Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	viper.SetEnvPrefix("MICROCART")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}
