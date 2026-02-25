package repository

import (
	"errors"
	"notification-service/errs"
	"notification-service/internal/models"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(sub *models.Subscription) error
	Delete(subs uint) error
	GetUsersByCategory(categoryID uint) ([]uint, error)
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(sub *models.Subscription) error {
	err:= r.db.Create(sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.ErrSubscriptionExists
		}
		return err
	}
	return nil
}

func (r *subscriptionRepository) Delete(subsID uint) error {
	result := r.db.Delete(&models.Subscription{}, subsID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrSubscriptionNotFound
	}
	return nil
}

func (r *subscriptionRepository) GetUsersByCategory(categoryID uint) ([]uint, error) {
	var userIDs []uint
	err := r.db.Model(&models.Subscription{}).
	Where("category_id = ?", categoryID).
	Pluck("user_id", &userIDs).Error
	if err != nil {
		if len(userIDs) == 0 {
			return nil, errs.ErrSubscriptionNotFound
		}
		return nil, err
	}
	return userIDs, nil
}

