package config

import (
	"github.com/spf13/viper"
)

var Clients []string

func Init() error {
	viper.SetConfigFile("config/config.ini")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("clients.names", &Clients); err != nil {
		return err
	}

	return nil
}