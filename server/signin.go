package server

import (
	"encoding/json"
	"errors"
	newErr "forum/internal/errorface"
	"forum/models"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//Sign - struct to get Signer datas
type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//SignIn - Sigin page
func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleSignInPage(w)
		return
	}
	if r.Method == http.MethodPost {
		s.handleSignIn(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleSignInPage - if SigIn GET method
func (s *Server) handleSignInPage(w http.ResponseWriter) {
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
	}
}

//handleSignIn - if SignIn POST method
func (s *Server) handleSignIn(w http.ResponseWriter, r *http.Request) {
	var signer Sign
	var user *models.User
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(bytes, &signer); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	user, err = s.store.GetUserByLogin(signer.Login)
	if err != nil {
		if errors.Is(err, newErr.ErrWrongLogin) {
			SendNotify(w, http.StatusBadRequest, newErr.ErrWrongLogin)
			return
		} else {
			logger(w, http.StatusBadRequest, err)
			return
		}
	}

	if err = comparer(&user.Password, signer.Password); err != nil {
		if errors.Is(err, newErr.ErrWrongPass) {
			SendNotify(w, http.StatusBadRequest, err)
			return
		} else {
			logger(w, http.StatusBadRequest, err)
			return
		}
	}

	cookie := s.cookiesStore.CreateSession(user.ID)
	http.SetCookie(w, cookie)
}

//comparer - check length of pass and compare pass & hash
func comparer(hash *string, pass string) error {
	if len(pass) < 8 || len(pass) > 32 {
		return newErr.ErrPassCompare
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(pass)); err != nil {
		return newErr.ErrWrongPass
	}
	return nil
}
