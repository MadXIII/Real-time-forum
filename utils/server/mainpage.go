package server

import (
	"encoding/json"
	newErr "forum/utils/internal/error"
	"net/http"
)

//MainPage ...
func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			logger(w, http.StatusNotFound, newErr.ErrNotFound)
			return
		}
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		posts, err := s.store.GetAllPosts()
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		bytes, err := json.Marshal(posts)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		w.Write(bytes)
		return
		//made get all posts
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
