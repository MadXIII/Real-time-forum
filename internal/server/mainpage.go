package server

import (
	"encoding/json"
	"fmt"
	"forum/internal/models"
	"io/ioutil"
	"net/http"
)

// MainPage - main page for backend route "/"
func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleMainPageGet(w, r)
		return
	}
	if r.Method == http.MethodPost {
		s.handleMaingPagePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// global id for Categories
var globCategoryID = 0

// handleMainPageGet - if MainPaige GET method
func (s *Server) handleMainPageGet(w http.ResponseWriter, r *http.Request) {
	getCategories, err := s.store.GetCategories()
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	getPosts, err := s.store.GetAllPostsByCategoryID(globCategoryID)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, GetAllPosts: %w", err))
		return
	}

	usernames, err := s.store.GetAllUsernamesID()
	fmt.Println(usernames)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, GetAllUsernamesID: %w", err))
		return
	}

	response := struct {
		Categories []models.Categories
		Posts      []models.Post
		Usernames  []models.OnlineUsers
	}{
		Categories: getCategories,
		Posts:      getPosts,
		Usernames:  usernames,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, Marshal(posts): %w", err))
		return
	}

	w.Write(bytes)
}

// handleMaingPagePost - if MainPaige POST method
func (s *Server) handleMaingPagePost(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	var categoryID models.Categories

	if err = json.Unmarshal(bytes, &categoryID); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	globCategoryID = categoryID.ID

	posts, err := s.store.GetAllPostsByCategoryID(globCategoryID)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err = json.Marshal(posts)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
}
