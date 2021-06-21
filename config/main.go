package config

import "github.com/spf13/viper"

type Config struct {
    Database *Database
    Server   *Server
}

func New() *Config {
    viper.AddConfigPath(".")
    viper.SetConfigFile(".env")

    _ = viper.ReadInConfig()

    return &Config{
        Database: newDatabase(),
        Server:   newServer(),
    }
}
