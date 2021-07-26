package server

import (
    "context"
    "github.com/killtw/hercules/config"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "gorm.io/gorm"
    "net/http"
    "os"
    "os/signal"
    "time"
)

type server struct {
    config *config.Config
    db     *gorm.DB
    echo   *echo.Echo
}

func New(config *config.Config, db *gorm.DB) *server {
    e := echo.New()

    return &server{config, db, e}
}

func (s server) Run() {
    if s.config.Server.UID {
        s.echo.Use(s.setUserMiddleware)
    }
    s.echo.Use(s.setDbMiddleware)
    s.echo.Use(middleware.Gzip())
    s.echo.Use(middleware.Logger())
    s.echo.Use(middleware.Recover())

    s.echo.HideBanner = true
    s.echo.Renderer = LayoutTemplate

    s.echo.Debug = s.config.Debug

    s.echo.Static("/js", "public/js")
    s.echo.Static("/images", "public/images")
    s.echo.Static("/css", "public/css")

    s.routerConfigure(e)

    go func() {
        if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
            s.echo.Logger.Fatal(err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    s.echo.Logger.Info("Shutting down the server")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := s.echo.Shutdown(ctx); err != nil {
        s.echo.Logger.Fatal("Server forced to shutdown:", err)
    }
}
