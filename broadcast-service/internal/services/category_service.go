package services

import (
	"broadcast-service/internal/models"
	"broadcast-service/internal/repository"
	"log/slog"
)

type CategoryService struct {
	logger *slog.Logger
	categoryRepo *repository.CategoryRepo
}

func NewCategoryService(categoryRepo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) List() ([]models.Category, error) {
	tickets, err := s.categoryRepo.List()
	if err != nil {
		s.logger.Error(err.Error())
		return tickets, err
	}

	return tickets, nil
}
