package repository

import (
	"broadcast-service/internal/models"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) List() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Model(&models.Category{}).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepo) GetById(id uint) (*models.Category, error) {
	var category models.Category

	if err := r.db.Model(&models.Category{}).Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
