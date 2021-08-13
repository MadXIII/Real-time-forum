package main

import (
	"fmt"
	"forum/database"
	"forum/database/sqlite"
	"log"
)

type Conf struct {
	Store database.Repository
}

func main() {
	store, err := sqlite.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	conf := Conf{Store: store}
	fmt.Println(conf)
}
