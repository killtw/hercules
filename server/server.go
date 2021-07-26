package server

import (
    "context"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "gorm.io/gorm"
    "github.com/killtw/hercules/config"
    "net/http"
    "os"
    "os/signal"
    "time"
)

type server struct {
    config *config.Server
    db     *gorm.DB
}

func New(config *config.Server, db *gorm.DB) *server {
    return &server{config, db}
}

func (s server) Run() {
    e := echo.New()

    if s.config.UID {
        e.Use(s.setUserMiddleware)
    }
    e.Use(s.setDbMiddleware)
    e.Use(middleware.Gzip())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.HideBanner = true
    e.Renderer = LayoutTemplate
    e.Debug = true

    e.Static("/js", "public/js")
    e.Static("/images", "public/images")
    e.Static("/css", "public/css")

    s.routerConfigure(e)

    go func() {
        if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
            e.Logger.Fatal(err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    e.Logger.Info("Shutting down the server")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := e.Shutdown(ctx); err != nil {
        e.Logger.Fatal("Server forced to shutdown:", err)
    }
}
