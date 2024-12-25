package model

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;unique;not null"`
	Username     string `gorm:"size:100;unique;not null"`
	PasswordHash string `gorm:"size:512;unique;not null"`
}
