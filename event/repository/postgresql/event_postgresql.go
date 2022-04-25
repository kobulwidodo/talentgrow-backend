package postgresql

import (
	"talentgrow-backend/domain"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) domain.EventRepository {
	return &EventRepository{db}
}

func (r *EventRepository) Create(event domain.Event) error {
	if err := r.db.Create(&event).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) FindAll() ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.Find(&events).Error; err != nil {
		return events, err
	}
	return events, nil
}

func (r *EventRepository) FindOne(id uint) (domain.Event, error) {
	var event domain.Event
	if err := r.db.Where("id = ?", id).First(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}

func (r *EventRepository) FindByType(key string) ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.Where("type = ?", key).Find(&events).Error; err != nil {
		return events, err
	}
	return events, nil
}

func (r *EventRepository) Update(event domain.Event) error {
	if err := r.db.Save(&event).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) Delete(event domain.Event) error {
	if err := r.db.Delete(&event).Error; err != nil {
		return err
	}
	return nil
}