package usecase

import "talentgrow-backend/domain"

type InternshipAppilcantUseCase struct {
	internshipApplicantPostgresRepository domain.InternshipApplicantPostgresRepository
	internshipRepository                  domain.InternshipRepository
}

func NewInternshipApplicantUseCase(iapr domain.InternshipApplicantPostgresRepository, ir domain.InternshipRepository) domain.InternshipAppilcantUseCase {
	return &InternshipAppilcantUseCase{internshipApplicantPostgresRepository: iapr, internshipRepository: ir}
}

func (u *InternshipAppilcantUseCase) Apply(input *domain.ApplyInternship) error {
	_, err := u.internshipRepository.FindOne(input.InternshipId)
	if err != nil {
		return err
	}
	internshipApplicant := domain.InternshipApplicant{
		UserId:       input.UserId,
		Age:          input.Age,
		College:      input.College,
		Major:        input.Major,
		CvLink:       input.CvLink,
		InternshipId: input.InternshipId,
	}
	if err := u.internshipApplicantPostgresRepository.Create(internshipApplicant); err != nil {
		return err
	}
	return nil
}
