package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func SetupRoutes(environment *echo.Echo) {

	renderer := &TemplateRenderer{
		template: template.Must(template.ParseGlob("templates/*.html")),
	}

	environment.Renderer = renderer

	environment.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]string{
			"Title":   "My aplicacion",
			"Heading": "!Hola, mundo!",
			"Message": "Bienvenido a mi aplicacion web con Echo y plantillas HTML.",
		})
	})

	environment.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World! it's a hi route")
	})

	environment.GET("/saludo/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello World! to "+name+"!")
	})

	environment.Static("/static", "static")
}
