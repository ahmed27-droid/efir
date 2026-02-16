package models

import "gorm.io/gorm"

type Role string

const (
	RoleReader Role = "reader"
	RoleAuthor Role = "author"
	RoleAdmin  Role = "admin"
)

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(28);uniqueIndex;not null"`
	Username     string `gorm:"type:varchar(20);not null"`
	PasswordHash string `gorm:"type:varchar(200);not null"`
	FirstName    string `gorm:"type:varchar(22);not null"`
	LastName     string `gorm:"type:varchar(22);not null"`
	Role         Role   `gorm:"type:varchar(20);default:'reader';not null"`
	IsActive     bool   `gorm:"default:true"`
}

type CreateUserRequest struct {
	Email     string	`json:"email" validate:"required,email"`
	Username  string	`json:"username" validate:"required,min=3,max=20"`
	Password  string	`json:"password" validate:"required,min=6,max=23"`
	FirstName string	`json:"first_name" validate:"required, min=4,max=14"`
	LastName  string	`json:"last_name" validate:"required,min=3,max=12"`
}

type UpdateUserRequest struct {
	FirstName string	`json:"email"`
	LastName  string	`json:"last_name"`
	Username  string	`json:"username"`
}
