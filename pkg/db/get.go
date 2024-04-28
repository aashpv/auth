package db

import (
	"AuthService/pkg/models"
	"database/sql"
	"errors"
	"log"
)

func (p *postgres) Get(email string) (user models.User, err error) {
	err = p.db.Get(&user, "SELECT email, password FROM users WHERE email = $1", email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user with email %s does not exist", email)
			return models.User{}, errors.New("user does not exist")
		}
		log.Printf("error occurred while getting user: %v", err)
		return models.User{}, err
	}

	log.Printf("retrieved user with email %s successfully", email)
	return user, nil
}
