package repository

import (
	"broadcast-service/internal/models"

	"gorm.io/gorm"
)

type BroadcastRepo struct {
	db *gorm.DB
}

func NewBroadcastRepo(db *gorm.DB) *BroadcastRepo {
	return &BroadcastRepo{db: db}
}

func (r *BroadcastRepo) Create(broadcast *models.Broadcast) error {
	return r.db.Create(broadcast).Error
}

func (r *BroadcastRepo) Start(id uint64) error {
	return r.db.Model(&models.Broadcast{}).
		Where("id = ?", id).
		Update("status", models.Live).
		Error
}

func (r *BroadcastRepo) GetById(id uint64) (*models.Broadcast, error) {
	var broadcast *models.Broadcast

	if err := r.db.Preload("Category").Model(&models.Broadcast{}).Where("id = ?", id).First(&broadcast).Error; err != nil {
		return nil, err
	}

	return broadcast, nil
}
