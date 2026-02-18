package dto

type RegisterRequest struct {
	Email     string	`json:"email" validate:"required,email"`
	Username  string	`json:"username" validate:"required,min=3,max=20"`
	Password  string	`json:"password" validate:"required,min=6,max=23"`
	FirstName string	`json:"first_name" validate:"required,min=4,max=14"`
	LastName  string	`json:"last_name" validate:"required,min=3,max=12"`
}

type UpdateUserRequest struct {
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Username  string	`json:"username"`
}


type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string `json:"token"`
}