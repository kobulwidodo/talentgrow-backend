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
