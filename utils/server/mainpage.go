package server

import (
	"encoding/json"
	"net/http"
)

//MainPage ...
func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
