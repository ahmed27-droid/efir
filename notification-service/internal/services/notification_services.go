package services

import (
	"fmt"
	"notification-service/internal/dto"
	"notification-service/internal/models"
	"notification-service/internal/repository"
)

type NotificationService interface {
	NotifyBroadcastStarted(event dto.BroadcastStartedEvent) error
	NotifyPost(event dto.PostCreatedEvent) error
	CreateNotifications(userIDs []uint, message string) error
	MarkAsRead(notificationID uint, userID uint) error
	MarkAllAsRead(userID uint) error
	GetUnreadCount(userID uint) (int64, error)
}

type notificationService struct {
	notifRepo repository.NotificationRepository
	subRepo   repository.SubscriptionRepository
}

func NewNotificationService(
	notifRepo repository.NotificationRepository,
	subRepo repository.SubscriptionRepository,
) NotificationService {
	return &notificationService{
		notifRepo: notifRepo,
		subRepo:   subRepo,
	}
}

func (s *notificationService) NotifyBroadcastStarted(event dto.BroadcastStartedEvent) error {
	users, err := s.subRepo.GetUsersByCategory(event.CategoryID)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Началась трансляция: %s", event.Title)

	return s.CreateNotifications(users, message)
}

func (s *notificationService) NotifyPost(event dto.PostCreatedEvent) error {
if event.Importance != "breaking" {
		return nil
	}

	users, err := s.subRepo.GetUsersByCategory(event.CategoryID)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Срочное обновление в трансляции: %s", event.Title)

	return s.CreateNotifications(users, message)
}

func (s *notificationService) CreateNotifications(userIDs []uint, message string) error {

	notifications := make([]models.Notification, 0, len(userIDs))

	for _, userID := range userIDs {
		notifications = append(notifications, models.Notification{
			UserID:  userID,
			Message: message,
			IsRead:  false,
		})
	}

	return s.notifRepo.CreateBatch(notifications)
}

func (s *notificationService) MarkAsRead(notificationID uint, userID uint) error{
	return s.notifRepo.MarkAsRead(notificationID, userID)
}

func (s *notificationService) MarkAllAsRead(userID uint) error {
	return s.notifRepo.MarkAllAsRead(userID)
}

func (s *notificationService) GetUnreadCount(userID uint) (int64, error) {
	return s.notifRepo.GetUnreadCount(userID)
}
