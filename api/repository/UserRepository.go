package repository

import (
	"go-todo/infrastructure"
	"go-todo/models"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) NewUser(user models.User) error {
	return u.db.DB.Create(&user).Error
}
