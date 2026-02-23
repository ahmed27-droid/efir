package dto

type ShowUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email,max=100"`
	Username  string `json:"username" binding:"required,min=3,max=20,alphanum"`
	Password  string `json:"password" binding:"required,min=8,max=64"`
	FirstName string `json:"first_name" binding:"required,min=2,max=50,alpha"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50,alpha"`
}

type UpdateUserRequest struct {
	UserID    *int    `json:"user_id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
