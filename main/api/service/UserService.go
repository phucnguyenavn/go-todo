package service

import (
	"go-todo/main/api/repository"
	"go-todo/main/models"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (u UserService) NewUser(user models.User) error {
	return u.userRepo.NewUser(user)
}
