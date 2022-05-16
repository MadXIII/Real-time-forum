package server

import (
	"fmt"
	"net/http"
)

// LogOut - logout user from session
func (s *Server) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ck, err := r.Cookie("session")
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, r.Cookie(session): %w", err))
			return
		}
		username, err := s.getUsernameByCookie(r)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, getUsernameByCookie: %w", err))
			return
		}

		logout(w, ck)
		if err := s.session.DeleteCookie(ck); err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, DeleteCookie: %w", err))
			return
		}

		s.session.SetOfflineUser(username)

		success(w, "Logout is Done")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
