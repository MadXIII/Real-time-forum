package server

import (
	"encoding/json"
	"fmt"
	"forum/utils/models"
	"io/ioutil"
	"net/http"
)

//MainPage ...
func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.handleMainPageGet(w, r)
		return
	}
	if r.Method == "POST" {
		s.handleMaingPagePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (s *Server) handleMainPageGet(w http.ResponseWriter, r *http.Request) {
	gotCategories, err := s.store.GetCategories()
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	gotPosts, err := s.store.GetAllPosts()
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, GetAllPosts: %w", err))
		return
	}

	response := struct {
		Categories []models.Categories
		Posts      []models.Post
	}{
		Categories: gotCategories,
		Posts:      gotPosts,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, Marshal(posts): %w", err))
		return
	}

	w.Write(bytes)
}

func (s *Server) handleMaingPagePost(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	var categoryID models.Categories

	if err = json.Unmarshal(bytes, &categoryID); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(categoryID.ID)

}
