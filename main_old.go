package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Echo/routes"
)

func main() {
	environment := echo.New()

	routes.SetupRoutes(environment)

	environment.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	environment.Start(":8080")
}
