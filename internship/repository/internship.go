package repository

import (
	"talentgrow-backend/domain"

	"gorm.io/gorm"
)

type InternshipRepository struct {
	db *gorm.DB
}

func NewInternshipRepository(db *gorm.DB) domain.InternshipRepository {
	return &InternshipRepository{db: db}
}

func (r *InternshipRepository) Create(internship domain.Internship) error {
	if err := r.db.Create(&internship).Error; err != nil {
		return err
	}

	return nil
}

func (r *InternshipRepository) FindOne(id uint) (domain.Internship, error) {
	var internship domain.Internship
	if err := r.db.Where("id = ?", id).First(&internship).Error; err != nil {
		return internship, err
	}

	return internship, nil
}

func (r *InternshipRepository) FindAll() ([]domain.Internship, error) {
	var internship []domain.Internship
	if err := r.db.Find(&internship).Error; err != nil {
		return internship, err
	}
	return internship, nil
}
