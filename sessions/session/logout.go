package session

import (
	"log"
	"net/http"
)

//LogOut - quit Session
func LogOut(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session")
	if err != nil {
		log.Println("test")
		log.Println(err)
	}
	ck.MaxAge = -1

	http.SetCookie(w, ck)
}
