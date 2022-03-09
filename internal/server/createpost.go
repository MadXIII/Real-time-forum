package server

import (
	"encoding/json"
	"fmt"
	newErr "forum/internal/error"
	"forum/internal/models"
	"io/ioutil"
	"net/http"
	"time"
)

// CreatePost - /newpost's handler
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleCreatePostPage(w, r)
		return
	}

	if r.Method == http.MethodPost {
		s.handleCreatePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

// handleCreatePostPage - if CreatPost GET method
func (s *Server) handleCreatePostPage(w http.ResponseWriter, r *http.Request) {
	categories, err := s.sqlStore.GetCategories()
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, GetCategories: %w", err))
		return
	}
	bytes, err := json.Marshal(&categories)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, Marshal: %w", err))
	}

	w.Write(bytes)
	return
}

// handleCreatePost - if CreatePost POST method
func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, ReadAll(r.Body): %w", err))
		return
	}

	if len(bytes) == 0 {
		logger(w, http.StatusBadRequest, newErr.ErrNilBody)
		return
	}
	var newPost models.Post
	if err = json.Unmarshal(bytes, &newPost); err != nil {
		logger(w, http.StatusBadRequest, fmt.Errorf("handleCreatePost, Unmarshal(newPost): %w", err))
		return
	}

	if err = s.checkNewPostDatas(&newPost); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	newPost.Username, err = s.getUsernameByCookie(r)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, getUsernameByCookie: %w", err))
		return
	}

	postID, err := s.sqlStore.InsertPost(&newPost)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, InsertPost: %w", err))
		return
	}

	// create object to Response about Success
	resp := struct {
		ID     int    `json:"id"`
		Notify string `json:"notify"`
	}{
		ID:     postID,
		Notify: "Post is Created",
	}

	bytes, err = json.Marshal(&resp)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
}

// checkNewPostDatas - check Post Datas before Insert it into db
func (s *Server) checkNewPostDatas(post *models.Post) error {
	if err := s.sqlStore.CheckCategoryID(post.CategoryID); err != nil {
		return err
	}

	if len(post.Title) == 0 || len(post.Title) > 32 {
		return newErr.ErrPostTitle
	}

	if len(post.Content) == 0 {
		return newErr.ErrPostContent
	}
	// Set date format
	post.Timestamp = time.Now().Format("2.Jan.2006, 15:04")
	return nil
}

// getUsernameByCookie - get Username from db, by GetIDByCookie
func (s *Server) getUsernameByCookie(req *http.Request) (string, error) {
	ck, err := req.Cookie("session")
	if err != nil {
		return "", fmt.Errorf("getUsernameByCookie, r.Cookie(\"session\"): %w", err)
	}

	id, err := s.cookiesStore.GetIDByCookie(ck)
	if err != nil {
		return "", err
	}

	username, err := s.sqlStore.GetUsernameByID(id)
	if err != nil {
		return "", err
	}

	return username, nil
}
