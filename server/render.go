package server

import (
    "errors"
    "github.com/labstack/echo/v4"
    "html/template"
    "io"
)

type TemplateRegistry struct {}

var LayoutTemplate = &TemplateRegistry{}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
    funcMap := template.FuncMap{
        "inc": func(i int) int {
            return i + 1
        },
    }

    tmpl, err := template.New(name).Funcs(funcMap).ParseFiles("views/" + name)
    if err != nil {
        err := errors.New("Template not found -> " + name)

        return err
    }

    return tmpl.Execute(w, data)
}

