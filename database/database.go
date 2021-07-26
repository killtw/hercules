package database

import (
    "fmt"
    "github.com/killtw/hercules/config"
    "gorm.io/driver/mysql"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

func New(c *config.Database) *gorm.DB {
    var driver gorm.Dialector

    switch c.Connection {
    case "sqlite":
        driver = sqlite.Open(c.Host)
        break
    case "mysql":
        dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", c.Username, c.Password, c.Host, c.Port, c.Database)
        driver = mysql.Open(dns)
        break
    default:
        log.Fatalln("No support drive")
    }

    db, err := gorm.Open(driver, &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    return db
}
