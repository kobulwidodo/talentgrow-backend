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
