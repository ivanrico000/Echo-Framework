package http

import (
	"Echo/internal/modules/rooms/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	service *application.RoomService
}

func NewRoomHandler(service *application.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

// --------------------------------------
// Create
// --------------------------------------
func (h *RoomHandler) Create(c echo.Context) error {
	var req application.RoomCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	room, err := h.service.CreateRoom(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, room)
}

// --------------------------------------
// Update
// --------------------------------------
func (h *RoomHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid room id"})
	}

	var req application.RoomUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	room, err := h.service.UpdateRoom(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, room)
}

// --------------------------------------
// Get By ID
// --------------------------------------
func (h *RoomHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	room, err := h.service.GetRoomById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, room)
}

// --------------------------------------
// Get By Number
// --------------------------------------
func (h *RoomHandler) GetByNumber(c echo.Context) error {
	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid number"})
	}

	room, err := h.service.GetRoomByNumber(number)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, room)
}

// --------------------------------------
// List
// --------------------------------------
func (h *RoomHandler) List(c echo.Context) error {
	rooms, err := h.service.ListRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, rooms)
}

// --------------------------------------
// Delete
// --------------------------------------
func (h *RoomHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	if err := h.service.DeleteRoom(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
