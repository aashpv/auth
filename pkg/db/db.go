package db

import (
	"github.com/aashpv/auth/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DataBase interface {
	Create(user models.User) (err error)
	Get(email string) (user models.User, err error)
}

type postgres struct {
	db *sqlx.DB
}

func New(dataSourceName string) (p DataBase, err error) {

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return nil, err
	}

	log.Println("database connection opened successfully")
	return &postgres{db: db}, nil
}
