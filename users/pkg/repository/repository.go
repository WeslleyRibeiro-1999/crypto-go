package repository

import (
	"fmt"
	"log"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserID(id int64) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int64) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		log.Printf("erro ao armazenar usuario no banco de dados: %v", err)
		return nil, fmt.Errorf("erro ao armazenar usuario no banco de dados: %v", err)
	}

	return user, nil
}

func (r *repository) GetUserID(id int64) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar usuario no banco de dados: %v", err)
	}
	return &user, nil
}

func (r *repository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar usuarios no banco de dados: %v", err)
	}

	return &users, nil
}

func (r *repository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("erro ao atualizar usuario no banco de dados: %v", err)
	}

	return user, nil
}

func (r *repository) DeleteUser(id int64) bool {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		log.Println("Erro ao deletar usuario")
		return false
	}

	return true
}
