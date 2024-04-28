package db

import (
	"AuthService/pkg/models"
	"errors"
	"log"
)

func (p *postgres) Create(user models.User) (err error) {
	_, err = p.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)",
		user.Email,
		user.Password,
	)
	if err != nil {
		log.Printf("error occurred while creating user: %v", err)
		return errors.New("user with such credentials already exist")
	}
	log.Println("user created successfully")
	return nil
}
