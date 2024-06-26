package db

import (
	"errors"
	"github.com/aashpv/auth/pkg/models"
	"log"
)

func (p *postgres) Create(user models.User) (err error) {
	_, err = p.db.Exec("INSERT INTO users (email, password, role, number) VALUES ($1, $2, $3, $4)",
		user.Email,
		user.Password,
		user.Role,
		user.Number,
	)
	if err != nil {
		log.Printf("error occurred while creating user: %v", err)
		return errors.New("user with such credentials already exist")
	}
	log.Println("user created successfully")
	return nil
}
