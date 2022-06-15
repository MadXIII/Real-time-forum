package http

import "net/http"

func (a *API) middleWare(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		if err := a.service.Check(ck.Value); err != nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		handler(w, r)
	}
}
