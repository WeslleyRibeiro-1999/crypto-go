package api

import (
	"net/http"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/models"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type HttpUser interface {
	CreateUser(c echo.Context) error
	GetUserID(c echo.Context) error
	GetAllUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type httpUser struct {
	usecase usecase.UsecaseUser
}

var _ HttpUser = (*httpUser)(nil)

func NewHandler(usecase usecase.UsecaseUser) HttpUser {
	return &httpUser{
		usecase: usecase,
	}
}

func (h *httpUser) CreateUser(c echo.Context) error {
	var req models.CreateUser

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.usecase.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *httpUser) GetUserID(c echo.Context) error {
	id := c.Param("id")

	user, err := h.usecase.GetUserbyID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *httpUser) GetAllUsers(c echo.Context) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *httpUser) UpdateUser(c echo.Context) error {
	var user models.UpdateUser

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := h.usecase.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *httpUser) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	ok, err := h.usecase.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error(), "deleted": false})
	}

	return c.JSON(http.StatusOK, map[string]bool{"deleted": ok})
}
