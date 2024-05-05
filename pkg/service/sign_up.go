package service

import (
	"errors"
	"github.com/aashpv/auth/pkg/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(user models.User) (err error) {
	if !IsValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	if !IsValidPassword(user.Password) {
		return errors.New("invalid password")
	}

	if !IsValidNumber(user.Number) {
		return errors.New("invalid phone number")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error occurred while hashing password: %v", err)
		return err
	}

	user.Password = string(hashedPass)
	user.Role = "USER" //всем новым пользователям выдается роль USER, на ADMIN меняется в БД

	err = s.pgs.Create(user)
	if err != nil {
		log.Printf("error occurred while creating user: %v", err)
		return err
	}

	log.Println("user signed up successfully")
	return nil
}
