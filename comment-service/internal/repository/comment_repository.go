package repository

import (
	"comment-service/internal/errs"
	"comment-service/internal/models"
	"errors"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(commentID uint) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(comID uint) error
	List(postID uint, page, limit int) ([]models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetByID(commentID uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.First(&comment, commentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrCommentNotFound
		}
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(comID uint) error {
	result := r.db.Delete(&models.Comment{}, comID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrCommentNotFound
	}
	return nil
}

func (r *commentRepository) List(postID uint, page, limit int) ([]models.Comment, error) {
	var comments []models.Comment

	offset := (page - 1) * limit

	if err := r.db.
		Where("post_id = ?", postID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
