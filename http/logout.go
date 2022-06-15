package http

import (
	"fmt"
	"net/http"
)

func (a *API) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ck, err := r.Cookie("session")
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, r.Cookie(session): %w", err))
			return
		}
		a.service.Logout(w, ck)
		if err := a.service.Cookie.Delete(ck); err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("LogOut, DeleteCookie: %w", err))
			return
		}
		success(w, "Logout is Done")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
