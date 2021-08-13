package server

import (
	"log"
	"net/http"
)

func RegHandler(w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		log.Println(w, "Register")
	})
	// switch r.Method {
	// case "POST":
	// 	//
	// case "GET":
	// 	//
	// default:
	// 	//
	// }
	// return nil
}
