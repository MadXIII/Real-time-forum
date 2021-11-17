package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Post struct {
	ID      string `json:"id"`
	Comment string `json:"commnet"`
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
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
	var postID Post

	if err := json.Unmarshal(bytes, &postID); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err := checkPostID(postID.ID); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	post, err := s.store.GetPostByID(postID.ID)
	if err != nil {
		logger(w, http.StatusBadRequest, err)
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

func checkPostID(id string) error {
	_, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return nil
}

// func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

// }
