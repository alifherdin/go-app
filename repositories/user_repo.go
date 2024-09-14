package repositories

import (
	"go-api/domains"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gdb *gorm.DB) *UserRepository {
	return &UserRepository{gdb}
}

func (ur *UserRepository) CreateUser(user *domains.User) error {
	return ur.db.Create(&user).Error
}
