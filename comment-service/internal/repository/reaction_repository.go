package repository

import (
	"commen-sService/internal/models"

	"gorm.io/gorm"
)

type ReactionRepository interface {
	Create(reaction *models.Reaction) error
	GetByID(reactionID uint) (*models.Reaction, error)
	Update(reaction *models.Reaction) error
	Delete(reaction *models.Reaction) error
	List(postID uint) (map[string]int64, error)
}

type reactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepository {
	return &reactionRepository{db: db}
}

func (r *reactionRepository) Create(reaction *models.Reaction) error{
	return  r.db.Create(reaction).Error
}

func (r *reactionRepository) GetByID(reactionID uint) (*models.Reaction, error){
	var reaction models.Reaction

	if err := r.db.First(&reaction, reactionID).Error; err != nil{
		return nil, err
	}
	return &reaction, nil
}

func (r *reactionRepository) Update(reaction *models.Reaction) error{
	return r.db.Save(reaction).Error
}

func (r *reactionRepository) Delete(reaction *models.Reaction) error{
	return r.db.Delete(reaction).Error
}

type reactionCountRaw struct{
	Type string
	Count int64
}
func (r *reactionRepository) List(postID uint) (map[string]int64,error){
	var rows []reactionCountRaw

	err := r.db.Model(&models.Reaction{}).
	Select("type, COUNT(*) as count").
	Where("postID = ?", postID).
	Group("type").
	Scan(&rows).Error

	if err != nil{
		return nil, err
	}

	result := make(map[string]int64)

	for _, row := range rows{
		result[row.Type]= row.Count
	}

	return result, nil


} 
