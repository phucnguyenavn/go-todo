package repository

import (
	"go-todo/infratructure"
	"go-todo/models"
)

type UserRepository struct {
	db infratructure.Database
}

func NewUserRepository(db infratructure.Database) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) NewUser(user models.User) error {
	return u.db.DB.Create(&user).Error
}
