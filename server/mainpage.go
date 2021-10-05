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
		if err = s.Parser(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = s.temp.Execute(w, nil)
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
