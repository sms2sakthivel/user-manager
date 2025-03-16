package model

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	ID       uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ucr *UserCreateRequest) GetDBObject() *User {
	return &User{Name: ucr.Name, Email: ucr.Email, Username: ucr.Username, PasswordHash: ucr.Password}
}

func (uur *UserUpdateRequest) GetDBObject() *User {
	return &User{ID: uur.ID, Name: uur.Name, Email: uur.Email, Username: uur.Username, PasswordHash: uur.Password}
}
