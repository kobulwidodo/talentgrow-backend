package usecase

import (
	"talentgrow-backend/domain"
)

type InternshipUseCase struct {
	internshipRepository domain.InternshipRepository
}

func NewInternshipUseCase(ir domain.InternshipRepository) domain.InternshipUseCase {
	return &InternshipUseCase{internshipRepository: ir}
}

func (u *InternshipUseCase) CreateInternship(input *domain.CreateInternship) error {
	internship := domain.Internship{
		Position:    input.Position,
		Description: input.Description,
		Company:     input.Company,
		IsPaid:      *input.IsPaid,
		UserId:      input.UserId,
	}
	err := u.internshipRepository.Create(internship)
	if err != nil {
		return err
	}
	return nil
}

func (u *InternshipUseCase) GetInternshipById(input *domain.FindInternship) (domain.Internship, error) {
	internship, err := u.internshipRepository.FindOne(input.Id)
	if err != nil {
		return internship, err
	}
	return internship, nil
}

func (u *InternshipUseCase) GetInternships() ([]domain.Internship, error) {
	internship, err := u.internshipRepository.FindAll()
	if err != nil {
		return internship, err
	}
	return internship, nil
}

func (u *InternshipUseCase) UpdateInternship(input *domain.UpdateInternship) error {
	internship, err := u.internshipRepository.FindOne(input.Id)
	if err != nil {
		return err
	}
	internship.Position = input.Position
	internship.Description = input.Description
	internship.Company = input.Company
	internship.IsPaid = *input.IsPaid

	if err := u.internshipRepository.Update(internship); err != nil {
		return err
	}
	return nil
}

func (u *InternshipUseCase) DeleteInternship(id uint) error {
	internship, err := u.internshipRepository.FindOne(id)
	if err != nil {
		return err
	}
	if err := u.internshipRepository.Delete(internship); err != nil {
		return err
	}
	return nil
}
