package repository

import (
	"commen-sService/internal/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(userID uint) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(comment *models.Comment) error
	List(postID uint) ([]models.Comment, error)
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

func (r *commentRepository) GetByID(userID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(comment, userID).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}

func (r *commentRepository) List(postID uint) ([]models.Comment, error){
	var comments []models.Comment
	if err :=  r.db.Find(comments, postID).Error; err != nil{
		return  nil, err
	}
	return  comments, nil
}
