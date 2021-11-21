package server

import (
	"net/http"
)

//middleWare - redirect if client not in session to SigninPage
func (s *Server) middleWare(login bool, handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if login {
			ck, err := r.Cookie("session")
			if err != nil {
				// logger(w, http.StatusInternalServerError, fmt.Errorf("middleWare, r.Cookie(session): %w", err))
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
				return
			}
			if err := s.cookiesStore.CheckCookie(ck.Value); err != nil {
				// logger(w, http.StatusInternalServerError, fmt.Errorf("middleWare, CheckCookie: %w", err))
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
				return
			}
		}
		handler(w, r)
		return
	}
}
