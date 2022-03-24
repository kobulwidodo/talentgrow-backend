package domain

import "gorm.io/gorm"

type EventParticipat struct {
	gorm.Model
	UserId  uint
	User    User
	EventId uint
	Event   Event
}
