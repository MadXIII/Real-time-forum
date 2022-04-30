package server

import (
	"fmt"
	newErr "forum/internal/error"
	"net/http"
)

// Index - main route "/" for client
func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := checkURLPath(r.URL.Path); err != nil {
			logger(w, http.StatusNotFound, err)
			return
		}
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("Index, Execute: %w", err))
			return
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// checkURLPath - check URL's
func checkURLPath(path string) error {
	if path != "/" && path != "/signin" && path != "/signup" && path != "/newpost" && path != "/post" && path != "/logout" && path != "/chat" {
		return newErr.ErrNotFound
	}
	return nil
}
