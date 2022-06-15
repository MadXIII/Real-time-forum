package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/madxiii/real-time-forum/model"
)

// SignIn - Sigin page
func (a *API) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleSignIn, ReadAll(r.Body): %w", err))
			return
		}

		var signer model.Sign
		if err = json.Unmarshal(bytes, &signer); err != nil {
			logger(w, http.StatusBadRequest, fmt.Errorf("handleSignIn, Unmarshal %w", err))
			return
		}
		id, status, err := a.service.Login(signer)
		if err != nil {
			logger(w, status, err)
			return
		}
		cookie, err := a.service.Create(id)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		http.SetCookie(w, cookie)
		success(w, "Login is Done")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
