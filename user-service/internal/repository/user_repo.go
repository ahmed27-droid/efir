package repository

import (
	"errors"
	"fmt"
	"user/internal/errs"
	"user/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(User *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByUsername(username string) (bool, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil

}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {

		return false, fmt.Errorf("%w: %v", errs.ErrCheckEmailExists, err)
	}

	return count > 0, nil
}

func (r *userRepository) ExistsByUsername(username string) (bool, error) {
	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error

	if err != nil {

		return false, fmt.Errorf("%w: %v", errs.ErrCheckUsernameExists, err)
	}
	

	return count > 0, nil
}


func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}