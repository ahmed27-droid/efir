package repository

import (
	"notification-service/errs"
	"notification-service/internal/models"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateBatch(notifications []models.Notification) error
	GetByUser(userID uint, limit, page int) ([]models.Notification, error)
	MarkAsRead(notificationID uint, userID uint) error
	MarkAllAsRead(userID uint) error
	GetUnreadCount(userID uint) (int64, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateBatch(notifications []models.Notification) error {
	if len(notifications) == 0 {
		return nil
	}
	return r.db.CreateInBatches(notifications, 500).Error
}

func (r *notificationRepository) GetByUser(userID uint, limit, page int) ([]models.Notification, error) {
	var notifications []models.Notification

	offset := (page - 1) * limit

	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) MarkAsRead(notificationID uint, userID uint) error {
	result := r.db.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrNotificationNotFound
	}
	return nil
}

func (r *notificationRepository) MarkAllAsRead(userID uint) error {
	result := r.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotificationNotFound
	}

	return nil
}

func (r *notificationRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64

	err := r.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}