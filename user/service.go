package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckUsername(input CheckUsername) (bool, error)
	GetUserByID(ID int) (User, error)
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
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found On That Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckUsername(input CheckUsername) (bool, error) {
	username := input.Username

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found with That ID")
	}

	return user, nil

}
