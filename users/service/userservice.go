package service

import (
	"github.com/sms2sakthivel/user-manager/users/model"
	"github.com/sms2sakthivel/user-manager/users/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

// GetUsers Gets all the users details
func (s *UserService) GetUsers() ([]model.UserResponse, error) {
	var users []model.UserResponse = []model.UserResponse{}
	dbUsers, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range dbUsers {
		users = append(users, *user.GetAPIResponseObject())
	}

	return users, nil
}

func (s *UserService) GetUser(id uint) (*model.UserResponse, error) {
	dbUser, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return dbUser.GetAPIResponseObject(), nil
}

func (s *UserService) CreateUser(userReq *model.UserCreateRequest) (*model.UserResponse, error) {
	user := userReq.GetDBObject()
	err := s.Repo.CreateUser(user)
	return user.GetAPIResponseObject(), err
}

func (s *UserService) UpdateUser(userReq *model.UserUpdateRequest) (*model.UserResponse, error) {
	user := userReq.GetDBObject()
	err := s.Repo.UpdateUser(user)
	return user.GetAPIResponseObject(), err
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}
