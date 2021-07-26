package config

import "github.com/spf13/viper"

type Database struct {
    Connection string
    Host       string
    Port       string
    Database   string
    Username   string
    Password   string
}

func newDatabase() *Database {
    viper.SetDefault("DB_CONNECTION", "sqlite")
    viper.SetDefault("DB_HOST", "database/database.sqlite")
    viper.SetDefault("DB_PORT", "")
    viper.SetDefault("DB_DATABASE", "")
    viper.SetDefault("DB_USERNAME", "")
    viper.SetDefault("DB_PASSWORD", "")

    return &Database{
        Connection: viper.GetString("DB_CONNECTION"),
        Host:       viper.GetString("DB_HOST"),
        Port:       viper.GetString("DB_PORT"),
        Database:   viper.GetString("DB_DATABASE"),
        Username:   viper.GetString("DB_USERNAME"),
        Password:   viper.GetString("DB_PASSWORD"),
    }
}
