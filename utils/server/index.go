package server

import (
	newErr "forum/utils/internal/error"
	"net/http"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := checkURLPath(r.URL.Path); err != nil {
			logger(w, http.StatusNotFound, newErr.ErrNotFound)
			return
		}
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func checkURLPath(path string) error {
	if path != "/" && path != "/signin" && path != "/signup" && path != "/newpost" && path != "/post" && path != "/logout" {
		return newErr.ErrNotFound
	}
	return nil
}
