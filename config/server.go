package config

import "github.com/spf13/viper"

type Server struct {
    Port string
    UID  bool
}

func newServer() *Server {
    viper.SetDefault("HTTP_PORT", "8080")
    viper.SetDefault("HTTP_UID", false)

    return &Server{
        Port: viper.GetString("HTTP_PORT"),
        UID:  viper.GetBool("HTTP_UID"),
    }
}
