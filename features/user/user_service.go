package user

import (
	"Find-Backend/core/utils"

	"github.com/google/uuid"
)


type service struct {
	repo Repository
}

type Service interface {
	GetAllUsers() ([]User, error)
	GetUser(id string) (*User, error)
	CreateUser(username, password string) error
}


func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) GetAllUsers() ([]User, error) {
	users, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *service) GetUser(id string) (*User, error){
	filters := map[string]interface{}{"id": id,}
	user, err :=  service.repo.FindOne(filters)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *service) CreateUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password) 
	if err != nil {
		return err
	}

	user := User{
		ID: uuid.New().String(),
		Username: username,
		Password: hashedPassword,
	}
	if err := service.repo.SaveUser(&user); err != nil {
		return err
	}

	return nil
}