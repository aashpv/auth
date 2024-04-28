package service

import (
	"github.com/aashpv/auth/pkg/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(user models.User) (err error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error occurred while hashing password: %v", err)
		return err
	}

	user.Password = string(hashedPass)

	err = s.pgs.Create(user)
	if err != nil {
		log.Printf("error occurred while creating user: %v", err)
		return err
	}

	log.Println("user signed up successfully")
	return nil
}
