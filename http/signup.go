package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/madxiii/real-time-forum/model"
	"golang.org/x/crypto/bcrypt"
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

		if err := isCorrectDatasToSignUp(newUser); err != nil {
			logger(w, http.StatusBadRequest, err)
			return
		}

		bytes, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreateAccount, GenerateFromPassword: %w", err))
			return
		}
		newUser.Password = string(bytes)

		if err = s.insertUserDB(&newUser); err != nil {
			logger(w, http.StatusBadRequest, err)
			return
		}

		cookie := s.cookiesStore.CreateSession(newUser.ID)
		http.SetCookie(w, cookie)
		success(w, "User successfully created")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleCreateAccount - if SignUp POST method
func (a *API) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
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

	if err := a.service.CheckUserData(newUser); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	cookie := s.cookiesStore.CreateSession(newUser.ID)
	http.SetCookie(w, cookie)
	success(w, "User successfully created")
}
