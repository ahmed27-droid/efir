package repository

import (
	"comment-Service/internal/errs"
	"comment-Service/internal/models"
	"errors"

	"gorm.io/gorm"
)

type ReactionRepository interface {
	Create(reaction *models.Reaction) error
	GetByID(reactionID uint) (*models.Reaction, error)
	Update(reaction *models.Reaction) error
	Delete(reactionID uint) error
	List(postID uint) (map[string]int64, error)
}

type reactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepository {
	return &reactionRepository{db: db}
}

func (r *reactionRepository) Create(reaction *models.Reaction) error {
	return r.db.Create(reaction).Error
}

func (r *reactionRepository) GetByID(reactionID uint) (*models.Reaction, error) {
	var reaction models.Reaction

	err := r.db.First(&reaction, reactionID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrReactionNotFound
		}
		return nil, err
	}
	return &reaction, nil
}

func (r *reactionRepository) Update(reaction *models.Reaction) error {
	return r.db.Save(reaction).Error
}

func (r *reactionRepository) Delete(reactionID uint) error {
	result := r.db.Delete(&models.Reaction{}, reactionID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrReactionNotFound
	}
	return nil
}

type reactionCountRaw struct {
	Type  string
	Count int64
}

func (r *reactionRepository) List(postID uint) (map[string]int64, error) {
	var rows []reactionCountRaw

	err := r.db.Model(&models.Reaction{}).
		Select("type, COUNT(*) as count").
		Where("post_id = ?", postID).
		Group("type").
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	result := make(map[string]int64)

	for _, row := range rows {
		result[row.Type] = row.Count
	}

	return result, nil

}
