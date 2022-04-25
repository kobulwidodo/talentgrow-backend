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

func (r *InternshipApplicantPostgresRepository) Create(internshipApplicant domain.InternshipApplicant) (uint, error) {
	if err := r.db.Create(&internshipApplicant).Error; err != nil {
		return 0, err
	}
	return internshipApplicant.ID, nil
}

func (r *InternshipApplicantPostgresRepository) FindOne(id uint) (domain.InternshipApplicant, error) {
	var data domain.InternshipApplicant
	if err := r.db.Where("id = ?", id).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *InternshipApplicantPostgresRepository) FindByUserIdAndInternId(userId uint, internshipId uint) (domain.InternshipApplicant, error) {
	var data domain.InternshipApplicant
	if err := r.db.Where("user_id = ? AND internship_id = ?", userId, internshipId).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *InternshipApplicantPostgresRepository) UpdateCv(data domain.InternshipApplicant) error {
	if err := r.db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}