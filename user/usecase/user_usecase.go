package usecase

import (
	"errors"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{UserRepository: ur}
}

func (u *UserUseCase) SignUp(input *domain.UserSignUp) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user := domain.User{
		Name:       input.Name,
		Email:      input.Email,
		Password:   string(hash),
		Occupation: input.Occupation,
	}
	if err := u.UserRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) SignIn(input *domain.UserSignIn) (string, error) {
	user, err := u.UserRepository.GetByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("credential not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateToken(user.ID, user.IsAdmin, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) GetMe(email string) (domain.User, error) {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}
