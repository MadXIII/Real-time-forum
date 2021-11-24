package server

import (
	"encoding/json"
	"fmt"
	newErr "forum/utils/internal/error"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//Sign - struct to store Signer datas
type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//SignIn - Sigin page
func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handleSignIn(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleSignIn - if SignIn POST method
func (s *Server) handleSignIn(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleSignIn, ReadAll(r.Body): %w", err))
		return
	}

	var signer Sign
	if err = json.Unmarshal(bytes, &signer); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleSignIn, Unmarshal %w", err))
		return
	}

	if err = checkLoginDatas(&signer); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	//wrong error GetUserByLogin with newErr ========================================================
	user, err := s.store.GetUserByLogin(signer.Login)
	if err != nil {
		logger(w, http.StatusBadRequest, newErr.ErrWrongLogin)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password)); err != nil {
		logger(w, http.StatusBadRequest, newErr.ErrWrongPass)
		return
	}

	cookie := s.cookiesStore.CreateSession(user.ID)
	http.SetCookie(w, cookie)
	success(w, "Login is Done")
}

//checkLoginDatas - is empty or too long login datas
func checkLoginDatas(user *Sign) error {
	if len(user.Login) == 0 || len(user.Login) > 32 {
		return newErr.ErrLoginData
	}
	if len(user.Password) < 8 || len(user.Password) > 32 {
		return newErr.ErrPassData
	}
	return nil
}
