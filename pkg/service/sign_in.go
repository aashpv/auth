package service

import (
	"AuthService/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignIn(email, password string) (jwtToken string, err error) {
	fmt.Println(email)
	user, err := s.pgs.Get(email)
	if err != nil {
		if errors.Is(err, errors.New("user does not exist")) {
			log.Printf("user with email %s does not exist", email)
			return "", errors.New("email or password is incorrect")
		}
		log.Printf("error occurred while getting user: %v", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("incorrect password for user with email %s", email)
		return "", errors.New("email or password is incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	jwtToken, err = token.SignedString(models.JwtKey)
	if err != nil {
		log.Printf("error occurred while generating JWT token: %v", err)
		return "", err
	}

	log.Printf("user with email %s signed in successfully", email)
	return jwtToken, nil
}
