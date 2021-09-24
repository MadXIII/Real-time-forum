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
		temp := Parser()
		w.WriteHeader(http.StatusOK)
		err = temp.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}

		return
		//made logout
		//made create post
		//made get all posts

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 Method not allowed"))
}
