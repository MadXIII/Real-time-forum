package server

import (
	newErr "forum/internal/error"
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
		return
		//made create post
		//made get all posts
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
