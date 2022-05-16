package main

import (
	"fmt"
	db "forum/internal/database/sqlite"
	"forum/internal/server"
	"forum/internal/sessions/session"
	"log"
	"os"
)

func main() {
	db := db.Store{}
	err := db.NewDB("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	usersList, err := db.GetAllUsernamesID()
	if err != nil {
		log.Fatal(err)
	}

	sessionService := session.NewSessionStore(usersList)
	server := server.NewServer(&db, sessionService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383"
	}
	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatal(err)
	}
	fmt.Println(usersList)
}
