package repository

import (
	"github.com/sms2sakthivel/user-manager/users/model"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}
