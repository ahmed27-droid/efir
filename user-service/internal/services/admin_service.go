package services

import (
	"net/http"
	"user/internal/errs"
	"user/internal/repository"
)

const (
	postURL = ""
	reactiontsURL = ""
)

type AdminService interface {
	DeleteUser(userID uint) error
}

type adminService struct {
	userRepo repository.UserRepository
	сlient  *http.Client
}

func NewAdminService(userRepo repository.UserRepository) AdminService {
	return &adminService{userRepo: userRepo,}
}

func (s *adminService) DeleteUser(userID uint) error {
	if userID == 0 {
		return errs.ErrInvalidUserID
	}

	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	if err := s.userRepo.Delete(userID); err != nil {
		return err
	}

	return nil
}
