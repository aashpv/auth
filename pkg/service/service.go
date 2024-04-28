package service

import (
	"github.com/aashpv/auth/pkg/db"
	"github.com/aashpv/auth/pkg/models"
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
