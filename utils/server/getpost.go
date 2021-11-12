package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type post struct {
	id int `json:"id"`
}

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

	var postID post

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(bytes, &postID); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	// postID, err := ParsePostID(r)
	// if err != nil {
	// 	logger(w, http.StatusInternalServerError, err)
	// 	return
	// }

	post, err := s.store.GetPostByID(postID.id)
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
	return
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
