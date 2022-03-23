package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type UserRepository interface {
	Create(user User) error
	GetByEmail(email string) (User, error)
}

type UserUseCase interface {
	SignUp(input *UserSignUp) error
	SignIn(input *UserSignIn) (string, error)
}

type UserSignUp struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type UserSignIn struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
