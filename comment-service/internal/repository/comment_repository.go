package repository

import (
	"commen-sService/internal/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(commentID uint) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(comment *models.Comment) error
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
	if err := r.db.First(&comment, commentID).Error; err != nil {
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

func (r *commentRepository) List(postID uint, page, limit int) ([]models.Comment, error){
	var comments []models.Comment

	offset := (page - 1) * limit

	if err :=  r.db.
	Where("post_id = ?", postID).
	Order("created_at desc").
	Limit(limit).
	Offset(offset).
	Find(&comments).Error; err != nil{
		return  nil, err
	}
	return  comments, nil
}
