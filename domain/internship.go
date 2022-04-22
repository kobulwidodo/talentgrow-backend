package domain

import "gorm.io/gorm"

type Internship struct {
	gorm.Model
	Position    string `json:"position"`
	Description string `json:"description"`
	Company     string `json:"company"`
	IsPaid      bool `json:"is_paid"`
	UserId      uint `json:"-"`
}

type InternshipRepository interface {
	Create(internship Internship) error
	FindOne(id uint) (Internship, error)
	FindAll() ([]Internship, error)
	Update(internship Internship) error
	Delete(internship Internship) error
}

type InternshipUseCase interface {
	CreateInternship(input *CreateInternship) error
	GetInternshipById(input *FindInternship) (Internship, error)
	GetInternships() ([]Internship, error)
	UpdateInternship(input *UpdateInternship) error
	DeleteInternship(id uint) error
}

type CreateInternship struct {
	Position    string `binding:"required"`
	Description string `binding:"required"`
	Company     string `binding:"required"`
	IsPaid      *bool  `json:"is_paid" binding:"required"`
	UserId      uint
}

type UpdateInternship struct {
	Id          uint
	Position    string `binding:"required"`
	Description string `binding:"required"`
	Company     string `binding:"required"`
	IsPaid      *bool `json:"is_paid" binding:"required"`
	UserId      uint
}

type FindInternship struct {
	Id uint `uri:"id" binding:"required"`
}
