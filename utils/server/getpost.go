package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {

	log.Println("here")
	if r.Method == http.MethodGet {
		s.handlerGetPostPage(w, r)
		return
	}
	// if r.URL.Path != "/post" {}
	//
	// if r.Method == http.MethodPost {
	// 	s.handlerGetPost(w, r)
	// 	return
	// }

}

func (s *Server) handlerGetPostPage(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	_, ep := path.Split(r.URL.Path)
	fmt.Println(ep)

	post, err := s.store.GetPostByID(5)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(post)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
	log.Println(bytes)
}

// func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

// }
