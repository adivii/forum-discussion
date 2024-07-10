package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server
		Database *Database
	}

	Server struct {
		Port int
	}

	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	// Singleton Concept
	once.Do(func() {
		// Fetch configuration from config.yaml
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		// Used to fetch config from .env file
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		// Assign value from config.yaml into variable
		// For example Database.Driver will be assigned value from database>driver
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
