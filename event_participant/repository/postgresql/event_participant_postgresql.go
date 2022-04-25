package postgresql

import (
	"talentgrow-backend/domain"

	"gorm.io/gorm"
)

type EventParticipantRepository struct {
	db *gorm.DB
}

func NewEventParticipantRepository(db *gorm.DB) domain.EventParticipantRepository {
	return &EventParticipantRepository{db}
}

func (r *EventParticipantRepository) Create(ep domain.EventParticipant) error {
	if err := r.db.Create(&ep).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventParticipantRepository) GetAll() ([]domain.EventParticipant, error) {
	var ep []domain.EventParticipant
	if err := r.db.Find(&ep).Error; err != nil {
		return ep, err
	}
	return ep, nil
}

func (r *EventParticipantRepository) GetByEventId(id uint) ([]domain.EventParticipant, error) {
	var ep []domain.EventParticipant
	if err := r.db.Where("event_id = ?", id).Error; err != nil {
		return ep, err
	}
	return ep, nil
}

func (r *EventParticipantRepository) GetByEventIdAndUserId(userId uint, eventId uint) (domain.EventParticipant, error) {
	var ep domain.EventParticipant
	if err := r.db.Where("event_id = ? AND user_id = ?", eventId, userId).Find(&ep).Error; err != nil {
		return ep, err
	}
	return ep, nil
}