package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	Type        string
	Date        time.Time
	UserId      uint
}
