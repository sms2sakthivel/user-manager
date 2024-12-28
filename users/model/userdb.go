package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}

func (user *User) GetAPIResponseObject() *UserResponse {
	return &UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Username: user.Username}
}
