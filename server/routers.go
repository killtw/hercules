package server

import (
    "github.com/labstack/echo/v4"
    "github.com/killtw/hercules/handlers"
    "net/http"
)

func (s server) routerConfigure(r *echo.Echo) {
    r.GET("/", handlers.HomeHandler)

    r.GET("healthz", heartbeatHandler)
}

func heartbeatHandler(c echo.Context) error {
    return c.NoContent(http.StatusOK)
}
