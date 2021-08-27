package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if r.Method == "GET" {
		w.Write([]byte("TEST"))
	}
	if r.Method == "POST" {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(bytes, &newUser) //
		if err != nil {
			log.Fatalln(err)
		}
		err = s.store.InsertUser(newUser)
		if err != nil {
			log.Println(err)
		}
	}
}
