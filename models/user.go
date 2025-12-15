package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Phone    string  `json:"phone"`
	IsHidden bool    `json:"is_hidden"`
	Rating   float64 `json:"rating"`
}

type CreateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int64  `json:"user_id,omitempty"`
}

type UpdateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int64  `json:"user_id,omitempty"`
}

type GetUsersResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Count   int    `json:"count"`
	Users   []User `json:"users"`
}

type GetUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    []User `json:"user"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
