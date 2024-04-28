package server

import (
	"AuthService/pkg/server/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Run() error
}

type server struct {
	addr   string
	router router.Router
}

func New(addr string, router router.Router) Server {
	return &server{
		addr:   addr,
		router: router,
	}
}

func (s *server) Run() error {
	srv := http.Server{
		Addr:         ":" + s.addr,
		Handler:      s.router.NewRouter(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("error start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT)
	<-stop

	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutdown server: %v", err)
		return err
	}

	log.Println("successful shutdown server")
	return nil
}
