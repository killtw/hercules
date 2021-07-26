package config

import "github.com/spf13/viper"

type Config struct {
    Database *Database
    Server   *Server
    Debug    bool
}

func New() *Config {
    viper.AddConfigPath(".")
    viper.SetConfigFile(".env")

    viper.SetDefault("APP_DEBUG", false)

    _ = viper.ReadInConfig()

    return &Config{
        Database: newDatabase(),
        Server:   newServer(),
        Debug:    viper.GetBool("APP_DEBUG"),
    }
}
