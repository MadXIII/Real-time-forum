package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/madxiii/real-time-forum/model"
)

func (a *API) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreateAccount, ReadAll(r.Body): %w", err))
			return
		}

		var newUser model.User

		if err = json.Unmarshal(bytes, &newUser); err != nil {
			logger(w, http.StatusBadRequest, fmt.Errorf("handleCreateAccount, Unmarshal: %w", err))
			return
		}

		if status, err := a.service.Register(&newUser); err != nil {
			logger(w, status, err)
			return
		}

		cookie, err := a.service.Create(newUser.ID)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		http.SetCookie(w, cookie)
		success(w, "User successfully created")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
