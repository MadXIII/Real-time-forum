package server

import (
	"fmt"
	"net/http"
)

//middleWare - if you need to be in session but you don't redirect you to SigninPage
func (s *Server) middleWare(login bool, handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(login)
		if login {
			ck, err := r.Cookie("session")
			if err != nil {
				fmt.Println("ERR", err)
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
				return
			}
			if err := s.cookiesStore.CheckCookie(ck.Value); err != nil {
				fmt.Println("ERR CHECK", err)
				http.Redirect(w, r, "/signin", http.StatusSeeOther)
				return
			}
		}
		handler(w, r)
		return
	}
}
