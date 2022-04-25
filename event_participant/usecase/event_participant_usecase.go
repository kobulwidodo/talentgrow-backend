package usecase

import (
	"errors"
	"talentgrow-backend/domain"
)

type EventParticipantUsecase struct {
	eventParticipantRepository domain.EventParticipantRepository
}

func NewEventParticipantUsecase(epr domain.EventParticipantRepository) domain.EventParticipantUsecase {
	return &EventParticipantUsecase{eventParticipantRepository: epr}
}

func (u *EventParticipantUsecase) Create(input *domain.CreateEventParticipant) error {
	data, err := u.eventParticipantRepository.GetByEventIdAndUserId(input.UserId, input.EventId)
	if err != nil {
		return err
	}
	if data.ID != 0 {
		return errors.New("already registered")
	}
	ep := domain.EventParticipant{
		UserId: input.UserId,
		EventId: input.EventId,
		University: input.University,
	}
	if err := u.eventParticipantRepository.Create(ep); err != nil {
		return err
	}
	return nil
}

func (u *EventParticipantUsecase) CheckIsRegisterd(userId uint, input *domain.FindEventParticipantUri) (domain.EventParticipantResponse, error) {
	var dataResponse domain.EventParticipantResponse
	data, err := u.eventParticipantRepository.GetByEventIdAndUserId(userId, input.EventId)
	if err != nil {
		return dataResponse, err
	}
	dataResponse.IsRegistered = true
	dataResponse.University = data.University
	if data.ID != 0 {
		return dataResponse, nil
	}
	dataResponse.IsRegistered = false
	return dataResponse, nil
}