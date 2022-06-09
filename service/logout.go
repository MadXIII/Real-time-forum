package service

import "net/http"

type Logout struct{}

func NewLogout() *Logout {
	return &Logout{}
}

// logout set cookies max age to -1
func (l *Logout) Logout(w http.ResponseWriter, ck *http.Cookie) {
	ck.MaxAge = -1
	http.SetCookie(w, ck)
	w.WriteHeader(http.StatusOK)
}
