package server

import (
	"encoding/json"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
	"io/ioutil"
	"net/http"
)

//CreatePost ...
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleNewPostPage(w)
		return
	}
	if r.Method == http.MethodPost {
		s.handleCreatePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleNewPostPage - if CreatePost GET method
func (s *Server) handleNewPostPage(w http.ResponseWriter) {
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
	}
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
	// newPost := models.Post{Timestamp: time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")}

	if err = json.Unmarshal(bytes, &newPost); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err = checkNewPostDatas(newPost); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	//cant find how to error
	if newPost.UserID = s.cookiesStore.GetIDByCookie(r); newPost.UserID < 0 {
		logger(w, http.StatusInternalServerError, newErr.ErrNoCookie)
		return
	}
	//cant find how to error
	if err = s.store.InsertPost(newPost); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
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
