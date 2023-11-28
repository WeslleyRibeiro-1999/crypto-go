package usecase

import (
	"fmt"
	"log"
	"strconv"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/models"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/pkg/repository"
)

type UsecaseUser interface {
	CreateUser(user *models.CreateUser) (*models.CreatedUser, error)
	GetUserbyID(idString string) (*models.UserResponse, error)
	GetAllUsers() (*[]models.UserResponse, error)
	UpdateUser(user *models.UpdateUser) (*models.UserResponse, error)
	DeleteUser(idString string) (bool, error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repo repository.Repository) UsecaseUser {
	return &usecase{
		repository: repo,
	}
}

func (u *usecase) CreateUser(user *models.CreateUser) (*models.CreatedUser, error) {
	fmt.Println("estamos aqui111")

	newUser, err := u.repository.CreateUser(&models.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}

	response := &models.CreatedUser{
		Status: true,
		Name:   newUser.Name,
		Email:  newUser.Email,
	}

	return response, nil
}

func (u *usecase) GetUserbyID(idString string) (*models.UserResponse, error) {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println("erro ao converter a string para int64: ", err)
		return nil, err
	}

	user, err := u.repository.GetUserID(id)
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	return userResponse, nil
}

func (u *usecase) GetAllUsers() (*[]models.UserResponse, error) {
	var response []models.UserResponse
	users, err := u.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		response = append(response, models.UserResponse{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})
	}

	return &response, nil
}

func (u *usecase) UpdateUser(user *models.UpdateUser) (*models.UserResponse, error) {
	getUser, err := u.repository.GetUserID(user.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.repository.UpdateUser(&models.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   getUser.CreatedAt,
	})
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		Name:        updatedUser.Name,
		Email:       updatedUser.Email,
		PhoneNumber: updatedUser.PhoneNumber,
	}

	return userResponse, nil
}

func (u *usecase) DeleteUser(idString string) (bool, error) {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println("erro ao converter a string para int64: ", err)
		return false, err
	}

	_, err = u.repository.GetUserID(id)
	if err != nil {
		return false, err
	}

	ok := u.repository.DeleteUser(id)
	if !ok {
		return false, err
	}

	return true, nil
}
