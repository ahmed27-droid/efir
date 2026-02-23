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
	GetByID(id uint) (*models.User, error)
	UpdateProfile(UserID uint, req dto.UpdateUserRequest) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	 jwtManager *auth.JWTManager
}

func NewUserService(userRepo repository.UserRepository, jwtManager *auth.JWTManager) UserService {
	return &userService{userRepo: userRepo, jwtManager: jwtManager}
}

func (s *userService) GetByID(id uint) (*models.User, error) {

	if id == 0 {
		return nil, errs.ErrInvalidUserID
	}
	user, err := s.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Register(req dto.RegisterRequest) (*models.User, error) {

	email := strings.TrimSpace(strings.ToLower(req.Email))
	username := strings.TrimSpace(req.Username)
	firsname := strings.TrimSpace(req.FirstName)
	lastname := strings.TrimSpace(req.LastName)

	if err := auth.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	emailExists, err := s.userRepo.ExistsByEmail(email)

	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, errs.ErrEmailAlreadyExists
	}

	usernameExists, err := s.userRepo.ExistsByUsername(username)

	if err != nil {
		return nil, err
	}

	if usernameExists {
		return nil, errs.ErrUsernameAlreadyExists
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	// hashedPassword, err :=  s.jwtManager.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:     email,
		Username:  username,
		FirstName: firsname,
		LastName:  lastname,
		Password:  hashedPassword,
		Role:      models.RoleReader,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(req dto.LoginRequest) (string, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	if err := auth.CheckPassword(req.Password, user.Password); err != nil {
		return "", errs.ErrInvalidCredentials
	}

	token, err := s.jwtManager.GenerateAccessToken(user.ID, string(user.Role))

	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *userService) UpdateProfile(UserID uint, req dto.UpdateUserRequest) (*models.User, error) {

	user, err := s.userRepo.GetByID(UserID)

	if err != nil {
		return nil, err
	}

	if req.Username != nil && *req.Username != user.Username {
		exists, err := s.userRepo.ExistsByUsername(*req.Username)

		if err != nil {
			return nil, err
		}

		if exists {
			return nil, errs.ErrUsernameAlreadyExists
		}
		user.Username = *req.Username
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}

	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
