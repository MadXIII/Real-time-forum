package server

import (
	"log"
	"net/http"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	// var posts models.Post
	if r.Method == http.MethodGet {
		if err = s.Parser(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		s.temp.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method not allowed"))
}
