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
    layout := "base.html"
    layoutInter := ctx.Get("layout")
    if layoutInter != nil {
        layout = layoutInter.(string)
    }

    funcMap := template.FuncMap{
        "inc": func(i int) int {
            return i + 1
        },
    }

    tmpl, err := template.New(layout).Funcs(funcMap).ParseFiles("views/" + layout, "views/" + name)
    if err != nil {
        err := errors.New("Template not found -> " + name)

        return err
    }

    return tmpl.Execute(w, data)
}

