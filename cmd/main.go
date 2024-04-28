package main

import (
	"AuthService/pkg/db"
	"AuthService/pkg/server"
	"AuthService/pkg/server/router"
	"AuthService/pkg/server/router/handlers"
	"AuthService/pkg/service"
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
