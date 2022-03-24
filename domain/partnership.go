package domain

import "gorm.io/gorm"

type Partnership struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
	Pic         string
	UserId      uint
	User        User
}
