package services

import (
	"SE_School/models"
	"errors"
	"net/mail"
)

type UserRepository interface {
	Add(u models.User) error
	Get(email string) (*models.User, error)
}

type UserService struct {
	Repo UserRepository
}

func (service *UserService) AddUser(user models.User) error {
	exists, err := service.Repo.Get(user.Email)
	if err != nil {
		return err
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		return errors.New("email is not valid")
	}

	if exists != nil {
		return errors.New("user with such email already exists")
	}

	if err = service.Repo.Add(user); err != nil {
		return err
	}

	return nil
}

func (service *UserService) LoginUser(user models.User) error {
	userByEmail, err := service.Repo.Get(user.Email)

	if err != nil {
		return err
	}

	if userByEmail == nil {
		return errors.New("no user with specified email")
	}

	if userByEmail.Password != user.Password {
		return errors.New("incorrect password specified")
	}

	return nil
}
