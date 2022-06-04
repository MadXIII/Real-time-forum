package main

import (
	"forum/internal/repository/sqlite"
	"forum/internal/server"
	"forum/internal/sessions/session"
	"forum/repository"
	"forum/repository/sqlite"
	"log"
	"os"
)

func main() {
	db, err := sqlite.New("forum.db")
	if err != nil {
		log.Fatal("error initializing db: %s", err.Error())
	}
	defer db.Close()

	repos := repository.New(db)

	sessionService := session.New()
	server := server.Init(&store, sessionService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8282"
	}

	server.ListenAndServe(":" + port)
}
