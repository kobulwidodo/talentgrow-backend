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
	Create(internshipApplicant InternshipApplicant) (uint, error)
	FindOne(id uint) (InternshipApplicant, error)
	FindByUserIdAndInternId(userId uint, internshipId uint) (InternshipApplicant, error)
	UpdateCv(data InternshipApplicant) error
}

type InternshipAppilcantUseCase interface {
	Apply(input *ApplyInternship) (uint, error)
	FindOne(input *FindApplicant) (InternshipApplicant, error)
	CheckIsRegistered(userId uint, internshipId uint) (CheckRegisteredResponse, error)
	UploadCv(path string, id uint) error
}

type ApplyInternship struct {
	Age          int    `binding:"required"`
	College      string `binding:"required"`
	Major        string `binding:"required"`
	CvLink       string
	UserId       uint
	InternshipId uint
}

type CheckRegisteredResponse struct {
	Age int `json:"age"`
	College string `json:"college"`
	Major string `json:"major"`
	IsRegistered bool `json:"is_registered"`
}

type FindInternshipUri struct {
	InternshipId uint `uri:"internship_id" binding:"required"`
}

type FindApplicant struct {
	Id uint `uri:"id" binding:"required"`
}
