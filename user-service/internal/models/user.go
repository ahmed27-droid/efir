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
	Email        string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Username     string `gorm:"type:varchar(20);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:varchar(200);not null"`
	FirstName    string `gorm:"type:varchar(22);not null"`
	LastName     string `gorm:"type:varchar(22);not null"`
	Role         Role   `gorm:"type:varchar(20);default:'reader';not null"`
	IsActive     bool   `gorm:"default:true"`
}

