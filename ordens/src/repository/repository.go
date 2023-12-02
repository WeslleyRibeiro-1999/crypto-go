package repository

import (
	"log"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/models"
	"gorm.io/gorm"
)

type RepositoryOrder interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	FindAllByUserID(userID int64) (*[]models.Order, error)
	FindOneByID(id int64) (*models.Order, error)
	FindAllOrders() (*[]models.Order, error)
	DeleteOrder(id int64) error
}

type repository struct {
	db *gorm.DB
}

var _ RepositoryOrder = (*repository)(nil)

func NewRepository(db *gorm.DB) RepositoryOrder {
	return &repository{db}
}

func (r *repository) CreateOrder(order *models.Order) (*models.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		log.Println("src.repository.CreateOrder(): erro ao armazenar pedido")
		return nil, err
	}

	return order, nil
}

func (r *repository) FindAllByUserID(userID int64) (*[]models.Order, error) {
	var orders []models.Order
	if err := r.db.Find(&orders, "user_id = ?", userID).Error; err != nil {
		log.Println("src.repository.FindAllByUserID(): erro ao buscar pedidos por usuario")
		return nil, err
	}

	return &orders, nil
}

func (r *repository) FindOneByID(id int64) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		log.Println("src.repository.FindOneByID(): erro ao buscar pedido por usuario")
		return nil, err
	}

	return &order, nil
}

func (r *repository) FindAllOrders() (*[]models.Order, error) {
	var order []models.Order
	if err := r.db.Find(&order).Error; err != nil {
		log.Println("src.repository.FindAllOrders(): erro ao buscar todos os pedido")
		return nil, err
	}

	return &order, nil
}

func (r *repository) DeleteOrder(id int64) error {
	if err := r.db.Delete(&models.Order{}, id).Error; err != nil {
		log.Println("src.repository.DeleteOrder(): erro ao buscar todos os pedido")
		return err
	}

	return nil
}
