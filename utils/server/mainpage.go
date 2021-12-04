package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//MainPage ...
func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		posts, err := s.store.GetAllPosts()
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, GetAllPosts: %w", err))
			return
		}

		bytes, err := json.Marshal(posts)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("MainPage, Marshal(posts): %w", err))
			return
		}

		w.Write(bytes)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
