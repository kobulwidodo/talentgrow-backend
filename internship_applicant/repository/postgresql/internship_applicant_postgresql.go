package postgresql

import (
	"talentgrow-backend/domain"

	"gorm.io/gorm"
)

type InternshipApplicantPostgresRepository struct {
	db *gorm.DB
}

func NewInternshipApplicantPostgresRepository(db *gorm.DB) domain.InternshipApplicantPostgresRepository {
	return &InternshipApplicantPostgresRepository{db: db}
}

func (r *InternshipApplicantPostgresRepository) Create(internshipApplicant domain.InternshipApplicant) error {
	if err := r.db.Create(&internshipApplicant).Error; err != nil {
		return err
	}
	return nil
}
