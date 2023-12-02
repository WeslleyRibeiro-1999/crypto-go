package api

import (
	"net/http"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/models"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/usecase"
	"github.com/labstack/echo/v4"
)

type HttpOrder interface {
	CreateOrder(c echo.Context) error
	GetAllbyUserID(c echo.Context) error
	GetAllOrders(c echo.Context) error
	GetOrderByID(c echo.Context) error
	DeleteOrder(c echo.Context) error
}

type httpOrder struct {
	usecase usecase.Usecase
}

var _ HttpOrder = (*httpOrder)(nil)

func NewHandler(usecase usecase.Usecase) HttpOrder {
	return &httpOrder{
		usecase: usecase,
	}
}

func (h *httpOrder) CreateOrder(c echo.Context) error {
	var order models.CreateOrderRequest

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := h.usecase.CreateOrder(&order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *httpOrder) GetAllbyUserID(c echo.Context) error {
	userID := c.Param("user_id")

	res, err := h.usecase.GetAllbyUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *httpOrder) GetAllOrders(c echo.Context) error {
	res, err := h.usecase.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *httpOrder) GetOrderByID(c echo.Context) error {
	id := c.Param("id")

	res, err := h.usecase.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *httpOrder) DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"status": true})
}
