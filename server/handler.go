package server

import (
	"encoding/json"
	"forum/database/sqlite"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
)

// func RegHandler(w http.ResponseWriter, r *http.Request) {

// 	http.HandleFunc("/registration", Регистрация)
// }

var Peremennaya models.User
var store sqlite.Store

func Регистрация(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("TEST"))
	}
	if r.Method == "POST" {
		bytes, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(bytes, &Peremennaya) //
		if err != nil {
			log.Fatalln(err)
		}

	}
	// err = sqlite.InsertUser(Peremennaya, &store)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(Peremennaya)
}
