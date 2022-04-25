package domain

import "gorm.io/gorm"

type EventParticipant struct {
	gorm.Model
	UserId  uint
	User    User
	University string
	EventId uint
	Event   Event
}

type EventParticipantRepository interface {
	Create(ep EventParticipant) error
	GetAll() ([]EventParticipant, error)
	GetByEventId(id uint) ([]EventParticipant, error)
	GetByEventIdAndUserId(userId uint, eventId uint) (EventParticipant, error)
}

type EventParticipantUsecase interface {
	Create(input *CreateEventParticipant) error
	CheckIsRegisterd(userId uint, input *FindEventParticipantUri) (EventParticipantResponse, error)
}

type CreateEventParticipant struct {
	University string `binding:"required"`
	EventId uint
	UserId uint
}

type FindEventParticipantUri struct {
	EventId uint `uri:"event_id" binding:"required"`
}

type EventParticipantResponse struct {
	IsRegistered bool `json:"is_registered"`
	University string `json:"university"`
}


