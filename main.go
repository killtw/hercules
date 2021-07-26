package main

import (
    "github.com/killtw/hercules/config"
    "github.com/killtw/hercules/database"
    "github.com/killtw/hercules/handlers"
    "github.com/killtw/hercules/server"
    "github.com/labstack/echo/v4"
)

func main() {
    configure := config.New()
    db := database.New(configure.Database)

    server.New(configure, db).Routes(func(e *echo.Echo) {
        e.GET("/", handlers.HomeHandler)
    }).Run()
}
