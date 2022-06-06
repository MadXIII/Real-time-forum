package main

import (
	"log"

	"github.com/madxiii/real-time-forum/http"
	"github.com/madxiii/real-time-forum/repository"
	"github.com/madxiii/real-time-forum/repository/sqlite"
	"github.com/madxiii/real-time-forum/server"
	"github.com/madxiii/real-time-forum/service"
)

func main() {
	db, err := sqlite.New("forum.db")
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
	services := service.New(repo)
	handlers := http.NewAPI(services).InitRoutes()

	serv := new(server.Server)

	if err := serv.Run("localhost:8383", handlers); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
