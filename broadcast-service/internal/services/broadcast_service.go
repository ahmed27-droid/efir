package services

import (
	"broadcast-service/internal/dto"
	"broadcast-service/internal/kafka"
	kafka_events "broadcast-service/internal/kafka/events"
	"broadcast-service/internal/models"
	"broadcast-service/internal/repository"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

type BroadcastService struct {
	logger        *slog.Logger
	broadcastRepo *repository.BroadcastRepo
	categoryRepo  *repository.CategoryRepo
	kafkaProducer *kafka.Producer
}

func NewBroadcastService(
	logger *slog.Logger,
	broadcastRepo *repository.BroadcastRepo,
	categoryRepo *repository.CategoryRepo,
	kafkaProducer *kafka.Producer,
) *BroadcastService {
	return &BroadcastService{
		logger:        logger,
		broadcastRepo: broadcastRepo,
		categoryRepo:  categoryRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (s *BroadcastService) Create(dto *dto.CreateBroadcastRequest) (*models.Broadcast, error) {
	var broadcast *models.Broadcast

	_, err := s.categoryRepo.GetById(dto.CategoryID)
	if err != nil {
		return nil, errors.New("category is not find")
	}

	broadcast = &models.Broadcast{
		CategoryID: dto.CategoryID,
		Name:       dto.Name,
		Status:     models.Draft,
	}

	if err := s.broadcastRepo.Create(broadcast); err != nil {
		return nil, fmt.Errorf("error creating broadcast: %w", err)
	}

	return broadcast, nil
}

func (s *BroadcastService) Start(ctx context.Context, id uint64) error {

	if err := s.broadcastRepo.Start(id); err != nil {
		return fmt.Errorf("error starting broadcast: %w", err)
	}
	broadcast, err := s.broadcastRepo.GetById(id)
	if err != nil {
		return err
	}

	event := kafka_events.BroadcastStartedEvent{
		BroadcastID: id,
		Title: broadcast.Name,
		Category: broadcast.Category.Name,
	}

	if err := s.kafkaProducer.BroadcastStarted(ctx, event); err != nil {
		s.logger.Warn("kafka publish failed", "error", err.Error())
	}

	return nil
}
