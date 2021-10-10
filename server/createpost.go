package server

import (
	"log"
	"net/http"
)

//CreatePost ...
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	} else if r.Method == http.MethodPost {

		log.Println("Need to finish this Part")
	}
}
