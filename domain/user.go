package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `gorm:"unique" json:"email"`
	Password   string `json:"-"`
	Occupation string `json:"occupation"`
	IsAdmin    bool `json:"is_admin"`
}

type UserRepository interface {
	Create(user User) error
	GetByEmail(email string) (User, error)
}

type UserUseCase interface {
	SignUp(input *UserSignUp) error
	SignIn(input *UserSignIn) (string, error)
	GetMe(email string) (User, error)
}

type UserSignUp struct {
	Name       string `binding:"required"`
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Occupation string `binding:"required"`
}

type UserSignIn struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
