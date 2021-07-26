package server

import (
    "context"
    "github.com/labstack/echo/v4"
    "net/http"
    "time"
)

func (s server) setUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        uid := c.FormValue("uid")

        if uid != "" {
            cookie := new(http.Cookie)
            cookie.Name = "UID"
            cookie.Value = uid
            cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
            c.SetCookie(cookie)
        } else {
            cookie, err := c.Cookie("UID")

            if err == nil {
                uid = cookie.Value
            }
        }

        ctx := context.WithValue(c.Request().Context(), "uid", uid)
        c.SetRequest(c.Request().WithContext(ctx))

        return next(c)
    }
}

func (s server) setDbMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        ctx := context.WithValue(c.Request().Context(), "db", s.db)
        c.SetRequest(c.Request().WithContext(ctx))

        return next(c)
    }
}
