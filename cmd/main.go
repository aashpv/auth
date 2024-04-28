package main

import (
	"github.com/aashpv/auth/pkg/db"
	"github.com/aashpv/auth/pkg/server"
	"github.com/aashpv/auth/pkg/server/router"
	"github.com/aashpv/auth/pkg/server/router/handlers"
	"github.com/aashpv/auth/pkg/service"
	"log"
)

func main() {
	dataSourceName := "user=postgres password=postgres dbname=flowers host=localhost port=5432 sslmode=disable"

	p, err := db.New(dataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	log.Println("initialized database successfully")

	src := service.New(p)
	log.Println("initialized service successfully")

	hrs := handlers.New(src)
	log.Println("initialized handlers successfully")

	rts := router.New(hrs)
	log.Println("initialized router successfully")

	srv := server.New("8081", rts)
	log.Println("initialized server successfully")

	err = srv.Run()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
