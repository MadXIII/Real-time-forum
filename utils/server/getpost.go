package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Id int `json:"id"`
}

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handlerGetPostPage(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (s *Server) handlerGetPostPage(w http.ResponseWriter, r *http.Request) {

	var postID Post

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(bytes, &postID); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	post, err := s.store.GetPostByID(postID.Id)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err = json.Marshal(post)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
	return
}

// func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

// }
