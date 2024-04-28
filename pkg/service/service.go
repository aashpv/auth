package service

import (
	"AuthService/pkg/db"
	"AuthService/pkg/models"
)

type Service interface {
	SignUp(user models.User) (err error)
	SignIn(email, password string) (jwtToken string, err error)
}

type service struct {
	pgs db.DataBase
}

func New(postgres db.DataBase) Service {
	return &service{pgs: postgres}
}
