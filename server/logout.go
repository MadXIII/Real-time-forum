package server

import (
	"log"
	"net/http"
)

//LogOut ...
func (s *Server) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	} else if r.Method == http.MethodPost {
		ck, err := r.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if err := s.cookiesStore.DeleteCookie(ck); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		ck.MaxAge = -1
		http.SetCookie(w, ck)
		w.WriteHeader(200)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 method not allowed"))
	return
}
