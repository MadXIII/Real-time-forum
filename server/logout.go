package server

import (
	"net/http"
)

//LogOut ...
func (s *Server) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.Method == http.MethodPost {
		ck, err := r.Cookie("session")
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		if err := s.cookiesStore.DeleteCookie(ck); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		logout(w, ck)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 method not allowed"))
	return
}
