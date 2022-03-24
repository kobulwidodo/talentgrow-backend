package domain

import "gorm.io/gorm"

type Internship struct {
	gorm.Model
	Position    string
	Description string
	Company     string
	IsPaid      bool
	UserId      uint
}
