package main

import (
	db "forum/utils/database/sqlite"
	"forum/utils/server"
	"forum/utils/sessions/session"
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
		port = "8080"
	}
	log.Println("Server is listening:", port)
	server.ListenAndServe(":" + port)
}
