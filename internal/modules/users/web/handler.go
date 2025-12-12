package web

import (
	"Echo/internal/modules/users/service"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack/v5"
)

type UserHandler struct {
	UC *service.UserUseCases
}

func NewUserHandler(uc *service.UserUseCases) *UserHandler {
	return &UserHandler{UC: uc}
}

// -------------------------------------
// Utilitarios MsgPack
// -------------------------------------
func decodeMsgPack(c echo.Context, v interface{}) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(body, v)
}

func respondMsgPack(c echo.Context, status int, v interface{}) error {
	data, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}
	return c.Blob(status, "application/msgpack", data)
}

// -------------------------------------
// Handlers REST
// -------------------------------------

func (h *UserHandler) Create(c echo.Context) error {
	var dto service.CreateUserDTO
	if err := decodeMsgPack(c, &dto); err != nil {
		return c.String(http.StatusBadRequest, "invalid msgpack payload")
	}

	if err := h.UC.CreateUser(dto); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return respondMsgPack(c, http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *UserHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	var dto service.UpdateUserDTO
	if err := decodeMsgPack(c, &dto); err != nil {
		return c.String(http.StatusBadRequest, "invalid msgpack payload")
	}

	if err := h.UC.UpdateUser(id, dto); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return respondMsgPack(c, http.StatusOK, map[string]string{"status": "updated"})
}

func (h *UserHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	if err := h.UC.DeleteUser(id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return respondMsgPack(c, http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *UserHandler) ListByHotel(c echo.Context) error {
	hotelStr := c.Param("hotel_id")
	hotelID, err := strconv.Atoi(hotelStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid hotel id")
	}

	users, err := h.UC.ListByHotel(hotelID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return respondMsgPack(c, http.StatusOK, users)
}

func (h *UserHandler) GetById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	user, err := h.UC.GetByID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return respondMsgPack(c, http.StatusOK, user)
}
