package main

import (
	db "forum/internal/database/sqlite"
	"forum/internal/server"
	"forum/internal/sessions/session"
	"log"
	"os"
)

func main() {
	store := db.Store{}
	err := store.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	defer store.Close()

	sessionService := session.New()
	server := server.Init(&store, sessionService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8282"
	}

	server.ListenAndServe(":" + port)
}
