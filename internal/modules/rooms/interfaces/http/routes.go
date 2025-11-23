package http

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoomRoutes(e *echo.Echo, handler *RoomHandler) {

	g := e.Group("/rooms")

	g.POST("", handler.Create)
	g.PATCH("/:id", handler.Update)
	g.GET("", handler.List)
	g.GET("/:id", handler.GetById)
	g.GET("/number/:number", handler.GetByNumber)
	g.DELETE("/:id", handler.Delete)
}
