package main

import (
	db "forum/database/sqlite"
	"forum/server"
	"log"
	"os"
)

func main() {
	store := db.Store{}
	err := store.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	server := server.Init(&store)

	// fmt.Println(conf)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is listening:", port)
	server.ListenAndServe(":" + port)
}
