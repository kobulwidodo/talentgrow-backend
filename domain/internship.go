package domain

import "gorm.io/gorm"

type Internship struct {
	gorm.Model
	Position    string
	Description string
	Company     string
	IsPaid      bool
	UserId      uint `json:"-"`
}

type InternshipRepository interface {
	Create(internship Internship) error
	FindOne(id uint) (Internship, error)
	FindAll() ([]Internship, error)
}

type InternshipUseCase interface {
	CreateInternship(input *CreateInternship) error
	GetInternshipById(input *FindInternship) (Internship, error)
	GetInternships() ([]Internship, error)
}

type CreateInternship struct {
	Position    string `binding:"required"`
	Description string `binding:"required"`
	Company     string `binding:"required"`
	IsPaid      *bool  `json:"is_paid" binding:"required"`
	UserId      uint
}

type FindInternship struct {
	Id uint `uri:"id" binding:"required"`
}
