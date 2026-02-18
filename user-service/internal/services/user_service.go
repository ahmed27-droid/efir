package services

import (
	"user/internal/dto"
	"user/internal/models"
)


type UserService interface {
	Register(req dto.RegisterRequest) (*models.User, error)
	
}
