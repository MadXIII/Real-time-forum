package server

import (
	"log"
	"net/http"
)

func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
			return
		}
		s.Parser()
		w.WriteHeader(http.StatusOK)
		err = s.temp.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method not allowed"))
}
