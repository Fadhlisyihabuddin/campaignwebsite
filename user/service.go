package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = "user"

	newuser, err := s.repository.Save(user)
	if err != nil {
		return newuser, err
	}
	return newuser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
