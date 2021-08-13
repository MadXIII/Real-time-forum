package server

import (
	"forum/models"
	"net/http"
)

func RegHandler(w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/registration", registration)
}

var newUser models.User

func registration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newUser.Nickname = r.FormValue("nickname")
		newUser.Age = r.FormValue("age")
	case "GET":
		//
	default:
		//
	}
}
