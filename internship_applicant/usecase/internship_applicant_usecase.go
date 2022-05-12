package usecase

import (
	"errors"
	"talentgrow-backend/domain"
)

type InternshipAppilcantUseCase struct {
	internshipApplicantPostgresRepository domain.InternshipApplicantPostgresRepository
	internshipRepository                  domain.InternshipRepository
}

func NewInternshipApplicantUseCase(iapr domain.InternshipApplicantPostgresRepository, ir domain.InternshipRepository) domain.InternshipAppilcantUseCase {
	return &InternshipAppilcantUseCase{internshipApplicantPostgresRepository: iapr, internshipRepository: ir}
}

func (u *InternshipAppilcantUseCase) Apply(input *domain.ApplyInternship) (uint, error) {
	existData, err := u.internshipApplicantPostgresRepository.FindByUserIdAndInternId(input.UserId, input.InternshipId)
	if err != nil {
		return 0, err
	}
	if existData.ID != 0 {
		return 0, errors.New("already registered")
	}
	_, err = u.internshipRepository.FindOne(input.InternshipId)
	if err != nil {
		return 0, err
	}
	internshipApplicant := domain.InternshipApplicant{
		UserId:       input.UserId,
		Age:          input.Age,
		College:      input.College,
		Major:        input.Major,
		CvLink:       input.CvLink,
		InternshipId: input.InternshipId,
	}
	id, err := u.internshipApplicantPostgresRepository.Create(internshipApplicant)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *InternshipAppilcantUseCase) FindOne(input *domain.FindApplicant) (domain.InternshipApplicant, error) {
	data, err := u.internshipApplicantPostgresRepository.FindOne(input.Id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (u *InternshipAppilcantUseCase) CheckIsRegistered(userId uint, internshipId uint) (domain.CheckRegisteredResponse, error) {
	data, err := u.internshipApplicantPostgresRepository.FindByUserIdAndInternId(userId, internshipId)
	if err != nil {
		return domain.CheckRegisteredResponse{}, err
	}
	res := domain.CheckRegisteredResponse{
		Age:          data.Age,
		College:      data.College,
		Major:        data.Major,
		IsRegistered: false,
	}
	if data.ID != 0 {
		res.IsRegistered = true
	}
	return res, nil
}

func (u *InternshipAppilcantUseCase) UploadCv(path string, id uint) error {
	data, err := u.internshipApplicantPostgresRepository.FindOne(id)
	if err != nil {
		return err
	}
	data.CvLink = path
	if err := u.internshipApplicantPostgresRepository.UpdateCv(data); err != nil {
		return err
	}
	return nil
}
