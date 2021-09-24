package server

import (
	"log"
	"net/http"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	// var posts []models.Post
	if r.Method == http.MethodGet {
		temp := Parser()
		w.WriteHeader(http.StatusOK)
		err = temp.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method not allowed"))
}
