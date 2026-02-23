package services

import (
	"comment-Service/internal/cache"
	"comment-Service/internal/client"
	"comment-Service/internal/dto"
	"comment-Service/internal/errs"
	"comment-Service/internal/models"
	"comment-Service/internal/repository"
)

type CommentServices interface {
	CreateComment(postID uint, req *dto.CreateCommentDTO) (*models.Comment, error)
	GetCommentByID(comID uint) (*models.Comment, error)
	UpdateComment(comID uint, req *dto.UpdateCommentDTO) (*models.Comment, error)
	DeleteComment(comID uint) error
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

func (s *commentServices) CreateComment(postID uint, req *dto.CreateCommentDTO) (*models.Comment, error) {
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

	comment := &models.Comment{
		PostID:  postID,
		UserID:  req.UserID,
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

func (s *commentServices) DeleteComment(comID uint) error {
	return s.commentRepo.Delete(comID)
}

func (s *commentServices) ListComments(postID uint, page, limit int) ([]models.Comment, error) {
	exists, err := s.broadcast.PostExists(postID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errs.ErrPostNotFound
	}

	return s.commentRepo.List(postID, page, limit)
}
