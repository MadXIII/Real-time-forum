package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
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
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	postID, err := ParsePostID(r)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	post, err := s.store.GetPostByID(postID)
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
}

// func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

// }

func ParsePostID(r *http.Request) (int, error) {
	r.ParseForm()
	val := r.Form.Get("id")

	postID, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return postID, nil
}
