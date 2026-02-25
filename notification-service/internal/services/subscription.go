package services

import (
	"notification-service/internal/dto"
	"notification-service/internal/models"
	"notification-service/internal/repository"
)

type SubscriptionService interface {
	Subscribe(req dto.CreateSubscriptionDTO) error
	Unsubscribe(subsID uint) error
	GetSubscribers(categoryID uint) ([]uint, error)
}

type subscriptionService struct {
	subRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subRepo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{subRepo: subRepo}
}

func (s *subscriptionService) Subscribe(req dto.CreateSubscriptionDTO) error {
	sub := &models.Subscription{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
	}
	return  s.subRepo.Create(sub)
}

func (s *subscriptionService) Unsubscribe(subsID uint) error {
	return s.subRepo.Delete(subsID)
}

func (s *subscriptionService) GetSubscribers(categoryID uint) ([]uint, error) {
	userIDs, err := s.subRepo.GetUsersByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}
