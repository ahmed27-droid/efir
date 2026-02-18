package services

import (
	
	"strings"
	"user/internal/auth"
	"user/internal/dto"
	"user/internal/errs"
	"user/internal/models"
	"user/internal/repository"
)


type UserService interface {
	Register(req dto.RegisterRequest) (*models.User, error)
	Login(req dto.LoginRequest) (string, error)

} 

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}


func (s *userService) Register(req dto.RegisterRequest) (*models.User, error) {

	email := strings.TrimSpace(strings.ToLower(req.Email))
	username := strings.TrimSpace(req.Username)
	firsname := strings.TrimSpace(req.FirstName)
	lastname := strings.TrimSpace(req.LastName)

	if err := auth.ValidatePassword(req.Password); err != nil {
		return nil, errs.ErrInvalidPassword
	}

	emailExists, err := s.userRepo.ExistsByEmail(email)

	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, errs.ErrCheckEmailExists
	}

	usernameExists, err := s.userRepo.ExistsByUsername(username)

	if err != nil {
		return nil, err
	}

	if usernameExists {
		return nil, errs.ErrCheckUsernameExists
	}

	hashedPassword, err := auth.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email: email,
		Username: username,
		FirstName: firsname,
		LastName: lastname,
		Password: hashedPassword,
		Role: models.RoleReader,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s * userService) Login(req dto.LoginRequest) (string, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	if err := auth.CheckPassword(req.Password, user.Password); err != nil {
		return "", errs.ErrInvalidCredentials
	}

	token, err := auth.GenerateAccessToken(user.ID, string(user.Role))

	if err != nil {
		return "", err
	}

	return token, nil

}