package config

import "github.com/spf13/viper"

func SetDefaults() {
	// server
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8080")

	// database
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.username", "dbuser")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.name", "nyx")

}
