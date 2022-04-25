package usecase

import (
	"talentgrow-backend/domain"
	"time"
)

type EventUsecase struct {
	eventRepository domain.EventRepository
}

func NewEventRepository(er domain.EventRepository) domain.EventUseCase {
	return &EventUsecase{eventRepository: er}
}

func (u *EventUsecase) Create(input *domain.CreateEventDto) error {
	layoutFormat := "2006-01-02 15:04:05 MST"
	date, _ := time.Parse(layoutFormat, input.Date + " WIB")
	event := domain.Event{
		Title: input.Title,
		Description: input.Description,
		Type: input.Type,
		Date: &date,
		UserId: input.UserId,
	}
	if err := u.eventRepository.Create(event); err != nil {
		return err
	}
	return nil
}

func (u *EventUsecase) GetAll() ([]domain.Event, error) {
	events, err := u.eventRepository.FindAll()
	if err != nil {
		return events, err
	}
	return events, nil
}

func (u *EventUsecase) GetByType(key string) ([]domain.Event, error) {
	events, err := u.eventRepository.FindByType(key)
	if err != nil {
		return events, err
	}
	return events, nil
}

func (u *EventUsecase) GetById(input *domain.FindEventUri) (domain.Event, error) {
	event, err := u.eventRepository.FindOne(input.Id)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (u *EventUsecase) Update(input *domain.UpdateEventDto) error {
	event, err := u.eventRepository.FindOne(input.Id)
	if err != nil {
		return err
	}
	event.Title = input.Title
	event.Description = input.Description
	event.Type = input.Type
	event.Date = &input.Date
	if err := u.eventRepository.Update(event); err != nil {
		return err
	}
	return nil
}

func (u *EventUsecase) Delete(id uint) error {
	event, err := u.eventRepository.FindOne(id)
	if err != nil {
		return err
	}
	err = u.eventRepository.Delete(event)
	if err != nil {
		return err
	}
	return nil
}