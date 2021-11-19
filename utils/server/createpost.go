package server

import (
	"encoding/json"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	ID     int    `json:"id"`
	Notify string `json:"notify"`
}

//CreatePost ...
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handleCreatePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleCreatePost - if CreatePost POST method
func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if len(bytes) == 0 {
		logger(w, http.StatusBadRequest, newErr.ErrNilBody)
		return
	}
	var newPost models.Post

	if err = json.Unmarshal(bytes, &newPost); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err = checkNewPostDatas(newPost); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	//Set date format
	newPost.Timestamp = time.Now().Format("2.Jan.2006, 15:04")

	newPost.Username, err = s.getUsernameByCookie(r)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	postID, err := s.store.InsertPost(&newPost)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	//create object to Response about Success
	resp := Response{postID, "Post is Created"}

	bytes, err = json.Marshal(resp)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
}

func checkNewPostDatas(post models.Post) error {
	if len(post.Title) == 0 || len(post.Title) > 32 {
		return newErr.ErrPostTitle
	}
	if len(post.Content) == 0 {
		return newErr.ErrPostContent
	}
	return nil
}

func (s *Server) getUsernameByCookie(req *http.Request) (string, error) {
	id, err := s.cookiesStore.GetIDByCookie(req)
	if err != nil {
		return "", err
	}
	username, err := s.store.GetUsernameByID(id)
	if err != nil {
		return "", err
	}

	return username, err
}
