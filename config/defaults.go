package config

import "github.com/spf13/viper"

func SetDefaults() {
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "8080")
}
