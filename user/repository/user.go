package repository

import (
	"talentgrow-backend/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).Limit(1).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
