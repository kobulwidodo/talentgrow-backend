package domain

import "gorm.io/gorm"

type InternshipApplicant struct {
	gorm.Model
	UserId       uint
	User         User
	Age          int
	College      string
	Major        string
	CvLink       string
	InternshipId uint
	Internship   Internship
}

type InternshipApplicantPostgresRepository interface {
	Create(internshipApplicant InternshipApplicant) error
	FindOne(id uint) (InternshipApplicant, error)
}

type InternshipAppilcantUseCase interface {
	Apply(input *ApplyInternship) error
	FindOne(input *FindApplicant) (InternshipApplicant, error)
}

type ApplyInternship struct {
	Age          int    `binding:"requried"`
	College      string `binding:"requried"`
	Major        string `binding:"requried"`
	CvLink       string `binding:"requried"`
	UserId       uint
	InternshipId uint
}

type FindInternshipUri struct {
	InternshipId uint `uri:"internship_id" binding:"required"`
}

type FindApplicant struct {
	Id uint `uri:"id" binding:"required"`
}
