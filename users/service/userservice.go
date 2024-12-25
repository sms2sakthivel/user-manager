package service

import (
	"github.com/sms2sakthivel/user-manager/users/model"
	"github.com/sms2sakthivel/user-manager/users/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}
