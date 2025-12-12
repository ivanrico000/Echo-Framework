package web

import "github.com/labstack/echo/v4"

func RegisterUserRoutes(e *echo.Echo, handler *UserHandler) {

	g := e.Group("/users")

	g.POST("", handler.Create)
	g.PATCH("/:id", handler.Update)
	g.GET("/hotel/:hotel_id", handler.ListByHotel)
	g.GET("/:id", handler.GetById)
	g.DELETE("/:id", handler.Delete)
}
