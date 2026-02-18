package services

import (
	"commen-sService/internal/cache"
	"commen-sService/internal/client"
	"commen-sService/internal/dto"
	"commen-sService/internal/errors"
	"commen-sService/internal/models"
	"commen-sService/internal/repository"
)

type CommentServices interface {
	CreateComment(userID uint, req *dto.CreateCommentDTO) (*models.Comment, error)
	GetCommentByID(commentID uint) (*models.Comment, error)
	UpdateComment(commentID uint, req *dto.UpdateCommentDTO) (*models.Comment, error)
	DeleteComment(commentID uint) error
	ListComments(postID uint, page, limit int) ([]models.Comment, error)
}

type commentServices struct {
	commentRepo repository.CommentRepository
	cache       cache.BroadcastCache
	broadcast   client.BroadcastClient
}

func NewCommentServices(
	commentRepo repository.CommentRepository, 
	cache cache.BroadcastCache, 
	broadcast client.BroadcastClient,
	) CommentServices {
	return &commentServices{
		commentRepo: commentRepo,
		cache:       cache,
		broadcast:   broadcast,
	}
}

func (s *commentServices) CreateComment(userID uint, req *dto.CreateCommentDTO) (*models.Comment, error) {
	exists, err := s.broadcast.PostExists(req.PostID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.ErrPostNotFound
	}

	active, found := s.cache.IsActive(req.PostID)

	if !found {
		active, err = s.broadcast.IsActive(req.PostID)
		if err != nil {
			return nil, err
		}
	}

	if !active {
		return nil, errors.ErrBroadcastNotActive
	}

	comment := &models.Comment{
		PostID:  req.PostID,
		UserID:  userID,
		Content: req.Content,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentServices) GetCommentByID(commentID uint) (*models.Comment, error) {
	return s.commentRepo.GetByID(commentID)
}

func (s *commentServices) UpdateComment(commentID uint, req *dto.UpdateCommentDTO) (*models.Comment, error) {
	comment, err := s.commentRepo.GetByID(commentID)	
	if err != nil {
		return nil, err
	}

	if req.Content != nil {
		comment.Content = *req.Content
	}

	if err := s.commentRepo.Update(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentServices) DeleteComment(commentID uint) error {
	comment, err := s.commentRepo.GetByID(commentID)
	if err != nil {
		return err
	}
	return s.commentRepo.Delete(comment)
}

func (s *commentServices) ListComments(postID uint, page, limit int) ([]models.Comment, error) {	
	exists, err := s.broadcast.PostExists(postID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.ErrPostNotFound
	}

	return s.commentRepo.List(postID, page, limit)
}
