package server

import (
	"log"
	"net/http"
)

//middleWare - if you need to be in session but you don't redirect you to SigninPage
func (s *Server) middleWare(login bool, handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.Println("aaaa?")

		if !login {
			log.Println("bbbbb?")

			handler(w, r)
			return
		}
		ck, err := r.Cookie("session")
		if err != nil {
			log.Println("ccccc?")
			w.WriteHeader(500)
			// http.Redirect(w, r, "/signin", http.StatusSeeOther)
			log.Println(err)
			return
		}

		if err := s.cookiesStore.CheckCookie(ck.Value); err != nil {
			log.Println("dddddd?")
			w.WriteHeader(500)

			// http.Redirect(w, r, "/signin", http.StatusSeeOther)
			log.Println(err)
			return
		}
		log.Println("ya svobodeeen")
		handler(w, r)
		return
	}
}
