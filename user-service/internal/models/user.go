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
	Email     string `json:"email" gorm:"type:varchar(50);uniqueIndex;not null"`
	Username  string `json:"username" gorm:"type:varchar(20);uniqueIndex;not null"`
	Password  string `json:"-" gorm:"type:varchar(200);not null"`
	FirstName string `json:"first_name" gorm:"type:varchar(22);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(22);not null"`
	Role      Role   `json:"role" gorm:"type:varchar(20);default:'reader';not null"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}
