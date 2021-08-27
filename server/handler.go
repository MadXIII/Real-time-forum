package server

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
)

var newUser models.User

func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("TEST"))
	}
	if r.Method == "POST" {
		bytes, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(bytes, &newUser) //
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(newUser.ID)
	}
}
