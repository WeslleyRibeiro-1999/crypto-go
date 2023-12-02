package usecase

import (
	"fmt"
	"log"
	"strconv"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/models"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/repository"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/service"
)

type Usecase interface {
	CreateOrder(order *models.CreateOrderRequest) (*models.CreateOrderResponse, error)
	GetAllbyUserID(userID string) (*[]models.Order, error)
	GetAllOrders() (*[]models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	DeleteOrder(id string) error
}

type usecase struct {
	repo repository.RepositoryOrder
}

var _ Usecase = (*usecase)(nil)

func NewUsecase(repo repository.RepositoryOrder) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) CreateOrder(order *models.CreateOrderRequest) (*models.CreateOrderResponse, error) {
	_, err := service.GetUserFromGRPC(order.UserID)
	if err != nil {
		return nil, err
	}

	newOrder, err := u.repo.CreateOrder(&models.Order{
		UserID:    order.UserID,
		Pair:      order.Pair,
		Amount:    order.Amount,
		Direction: order.Direction,
		Type:      models.MarketOrder,
	})
	if err != nil {
		return nil, err
	}

	res := models.CreateOrderResponse{
		Pair:      newOrder.Pair,
		Amount:    newOrder.Amount,
		Direction: newOrder.Direction,
	}

	return &res, nil
}

func (u *usecase) GetAllbyUserID(userID string) (*[]models.Order, error) {
	user, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Println("src.usecase.GetAllbyUserID(): erro ao converter string para int")
		return nil, err
	}

	_, err = service.GetUserFromGRPC(user)
	if err != nil {
		return nil, err
	}

	orders, err := u.repo.FindAllByUserID(user)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (u *usecase) GetAllOrders() (*[]models.Order, error) {
	orders, err := u.repo.FindAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (u *usecase) GetOrderByID(id string) (*models.Order, error) {
	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	order, err := u.repo.FindOneByID(orderID)
	if err != nil {
		log.Println("src.usecase.GetOrderByID(): erro ao converter string para int")
		return nil, err
	}

	return order, err
}

func (u *usecase) DeleteOrder(id string) error {
	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("src.usecase.DeleteOrder(): erro ao converter string para int")
		return err
	}

	order, err := u.repo.FindOneByID(orderID)
	if err != nil {
		return err
	}

	if order.Type != "limit" {
		return fmt.Errorf("erro ao deletar, pedido nao e do tipo limit")
	}

	err = u.repo.DeleteOrder(orderID)
	if err != nil {
		return nil
	}

	return nil
}
