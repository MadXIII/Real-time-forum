package main

import (
	"fmt"
	"forum/database"
	"forum/database/sqlite"
	"forum/server"
	"log"
	"net/http"
	"os"
)

type Conf struct {
	Store database.Repository
}

func main() {
	store, err := sqlite.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/registration", server.Регистрация)

	conf := Conf{Store: store}
	fmt.Println(conf)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is listening:", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
