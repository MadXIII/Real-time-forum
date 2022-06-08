package main

import (
	"log"

	"github.com/madxiii/real-time-forum/config"
	"github.com/madxiii/real-time-forum/http"
	"github.com/madxiii/real-time-forum/repository"
	"github.com/madxiii/real-time-forum/repository/sqlite"
	"github.com/madxiii/real-time-forum/server"
	"github.com/madxiii/real-time-forum/service"
)

func main() {
	configs, err := config.Get("./config/config.yml")
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := sqlite.New(configs.DB)
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
	services := service.New(repo)
	handlers := http.NewAPI(services).InitRoutes()

	serv := new(server.Server)

	if err := serv.Run(configs.Port, handlers); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
