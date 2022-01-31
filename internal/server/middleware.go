package server

import (
	"net/http"
)

//middleWare - redirect if client not in session to SigninPage
func (s *Server) middleWare(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		if err := s.cookiesStore.CheckCookie(ck.Value); err != nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		handler(w, r)
		return
	}
}
