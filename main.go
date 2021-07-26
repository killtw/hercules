package main

import (
    "github.com/killtw/hercules/config"
    "github.com/killtw/hercules/database"
    "github.com/killtw/hercules/server"
)

func main() {
    configure := config.New()
    db := database.New(configure.Database)

    server.New(configure, db).Run()
}
