package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Date        *time.Time `json:"date"`
	UserId      uint `json:"-"`
}

type EventRepository interface {
	Create(event Event) error
	FindAll() ([]Event, error)
	FindOne(id uint) (Event, error)
	FindByType(key string) ([]Event, error)
	Update(event Event) error
	Delete(event Event) error
}

type EventUseCase interface {
	Create(input *CreateEventDto) error
	GetAll() ([]Event, error)
	GetByType(key string) ([]Event, error)
	GetById(input *FindEventUri) (Event, error)
	Update(input *UpdateEventDto) error
	Delete(id uint) error
}

type CreateEventDto struct {
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Type        string `binding:"required"`
	Date        string `binding:"required"`
	UserId uint
}

type FindEventUri struct {
	Id uint `uri:"id" binding:"required"`
}

type FindEventsType struct {
	Type string `uri:"type" binding:"required"`
}

type UpdateEventDto struct {
	Id uint
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Type        string `binding:"required"`
	Date        time.Time `binding:"required"`
}