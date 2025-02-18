package model

type UserCreateRequest struct {
	Email       string `json:"email"`
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type UserUpdateRequest struct {
	ID          uint   `json:"user_id"`
	Email       string `json:"email"`
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	ID          uint   `json:"user_id"`
	Email       string `json:"email"`
	Username    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
}

func (ucr *UserCreateRequest) GetDBObject() *User {
	return &User{Name: ucr.Name, Email: ucr.Email, Username: ucr.Username, PasswordHash: ucr.Password, PhoneNumber: ucr.PhoneNumber}
}

func (uur *UserUpdateRequest) GetDBObject() *User {
	return &User{ID: uur.ID, Name: uur.Name, Email: uur.Email, Username: uur.Username, PasswordHash: uur.Password, PhoneNumber: uur.PhoneNumber}
}
