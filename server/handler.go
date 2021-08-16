package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
)

// func RegHandler(w http.ResponseWriter, r *http.Request) {

// 	http.HandleFunc("/registration", Регистрация)
// }

var newUser models.User

func Регистрация(w http.ResponseWriter, r *http.Request) {
	var peremennaya models.User
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &peremennaya) //
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(peremennaya)
}
