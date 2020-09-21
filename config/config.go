package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

// New config
func New() *viper.Viper {
	if config == nil {
		env := os.Getenv("ENV")
		if env == "" {
			env = "default"
		}

		config = viper.New()
		config.SetConfigName(env)
		config.SetConfigType("yaml")
		config.AddConfigPath("./config")
		err := config.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	return config
}
