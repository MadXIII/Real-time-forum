package main

import (
	db "forum/database/sqlite"
	"forum/server"
	"log"
	"os"
)

func main() {
	var store db.Store
	err := store.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	server := server.Init(store)
	// conf := Conf{Store: store}
	// err = sqlite.InsertUser(models.User{
	// 	Age: 23,
	// }, store)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(conf)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is listening:", port)
	server.ListenAndServe(":" + port)

	// err = http.ListenAndServe(":"+port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
