package services

import (
	"comment-Service/internal/cache"
	"comment-Service/internal/client"
	"comment-Service/internal/dto"
	"comment-Service/internal/errs"
	"comment-Service/internal/models"
	"comment-Service/internal/repository"
)

type ReactionServices interface {
	CreateReaction(postID uint, req dto.CreateReactionDTO) (*models.Reaction, error)
	UpdateReaction(reactionID uint, req dto.UpdateReactionDTO) (*models.Reaction, error)
	DeleteReaction(reactionID uint) error
	ListReaction(postID uint) (map[string]int64, error)
}

type reactionService struct {
	reactionRepo repository.ReactionRepository
	cache        cache.BroadcastCache
	broadcast    client.BroadcastClient
}

func NewReactionService(
	reactionRepo repository.ReactionRepository,
	cache cache.BroadcastCache,
	broadcast client.BroadcastClient,
) ReactionServices {
	return &reactionService{
		reactionRepo: reactionRepo,
		cache:        cache,
		broadcast:    broadcast,
	}
}

func (s *reactionService) CreateReaction(postID uint, req dto.CreateReactionDTO) (*models.Reaction, error) {
	exists, err := s.broadcast.PostExists(postID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errs.ErrPostNotFound
	}

	active, found := s.cache.IsActive(postID)
	if !found {
		active, err = s.broadcast.IsActive(postID)
		if err != nil {
			return nil, err
		}
	}

	if !active {
		return nil, errs.ErrBroadcastNotActive
	}

	reaction := &models.Reaction{
		PostID: postID,
		UserID: req.UserID,
		Type:   models.ReactionType(req.Type),
	}

	if err := s.reactionRepo.Create(reaction); err != nil {
		return nil, err
	}

	return reaction, nil
}

func (s *reactionService) UpdateReaction(reactionID uint, req dto.UpdateReactionDTO) (*models.Reaction, error) {
	reaction, err := s.reactionRepo.GetByID(reactionID)
	if err != nil {
		return nil, err
	}

	if req.Type != nil {
		reaction.Type = models.ReactionType(*req.Type)
	}

	if err := s.reactionRepo.Update(reaction); err != nil {
		return nil, err
	}
	return reaction, nil
}

func (s *reactionService) DeleteReaction(reactionID uint) error {
	return s.reactionRepo.Delete(reactionID)
}

func (s *reactionService) ListReaction(postID uint) (map[string]int64, error) {
	exists, err := s.broadcast.PostExists(postID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errs.ErrPostNotFound
	}
	return s.reactionRepo.List(postID)
}
