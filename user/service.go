package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, filelocation string) (User, error)
	FindUserByID(id int) (User, error)
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

	user.PasswordHash = string(password)
	user.Role = "user"

	findEmail, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return findEmail, err
	}

	if user.Email == findEmail.Email {
		return findEmail, errors.New("email udah ada")
	}

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

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(id int, filelocation string) (User, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation

	userUpdate, err := s.repository.Update(user)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *service) FindUserByID(id int) (User, error) {

	user, err := s.repository.FindByID(id)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("userid not found")
	}

	return user, nil
}
