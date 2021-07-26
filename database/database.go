package database

import (
    "fmt"
    "github.com/killtw/hercules/config"
    "gorm.io/driver/mysql"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func New(c *config.Database) *gorm.DB {
    var driver gorm.Dialector

    if c.Connection == "mysql" {
        dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", c.Username, c.Password, c.Host, c.Port, c.Database)
        driver = mysql.Open(dns)
    } else {
        driver = sqlite.Open(c.Host)
    }

    db, err := gorm.Open(driver, &gorm.Config{})

    if err != nil {
        panic(err)
    }

    db.AutoMigrate()

    return db
}
