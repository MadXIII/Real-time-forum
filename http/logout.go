package http

import (
	"fmt"
	"net/http"
)

func (a *API) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ck, err := r.Cookie("session")
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, r.Cookie(session): %w", err))
			return
		}
		logout(w, ck)
		if err := a.service.Cookie.Delete(ck); err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, DeleteCookie: %w", err))
			return
		}
		success(w, "Logout is Done")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
