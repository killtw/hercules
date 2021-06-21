package main

import (
    "hercules/config"
    "hercules/database"
    "hercules/server"
)

func main() {
    configure := config.New()
    db := database.New(configure.Database)

    server.New(configure.Server, db).Run()
}
