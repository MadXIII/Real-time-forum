package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handlerGetPostPage(w)
		return
	}
	// if r.Method == http.MethodPost {
	// 	s.handlerGetPost(w, r)
	// 	return
	// }

}

func (s *Server) handlerGetPostPage(w http.ResponseWriter) {
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	post, err := s.store.GetPostByID(1)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
	}

	bytes, err := json.Marshal(post)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
	}

	w.Write(bytes)
	log.Println(bytes)

}

// func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

// }
