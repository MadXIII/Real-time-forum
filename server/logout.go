package server

import (
	"fmt"
	"net/http"
)

//LogOut ...
func (s *Server) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("1")
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			fmt.Println("2")
			logger(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.Method == http.MethodPost {
		fmt.Println("3")
		ck, err := r.Cookie("session")
		if err != nil {
			fmt.Println("4")
			logger(w, http.StatusInternalServerError, err)
			return
		}
		logout(w, ck)
		if err := s.cookiesStore.DeleteCookie(ck); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		// ck.MaxAge = -1
		// http.SetCookie(w, ck)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
