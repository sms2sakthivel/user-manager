package repository

import (
	"github.com/sms2sakthivel/user-manager/users/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (repo *GormUserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repo.DB.Find(&users).Error
	return users, err
}

func (repo *GormUserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

func (repo *GormUserRepository) CreateUser(user *model.User) error {
	return repo.DB.Create(user).Error
}

func (repo *GormUserRepository) UpdateUser(user *model.User) error {
	return repo.DB.Save(user).Error
}

func (repo *GormUserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&model.User{}, id).Error
}
