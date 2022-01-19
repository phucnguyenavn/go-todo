package repository

import (
	"go-todo/main/infratructure"
	"go-todo/main/models"
)

type UserRepository struct {
	db infratructure.Database
}

func NewUserRepository(db infratructure.Database) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) NewUser(user models.User) error{
	return u.db.DB.Create(&user).Error
}
