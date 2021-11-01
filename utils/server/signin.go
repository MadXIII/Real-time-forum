package server

import (
	"encoding/json"
	newErr "forum/utils/internal/error"
	"io/ioutil"
	"log"
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
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	var signer Sign
	if err = json.Unmarshal(bytes, &signer); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err = checkLoginDatas(&signer); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	user, err := s.store.GetUserByLogin(signer.Login)
	if err != nil {
		logger(w, http.StatusBadRequest, newErr.ErrWrongLogin)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password)); err != nil {
		log.Println(user.Password, signer.Password)
		logger(w, http.StatusBadRequest, err)
		return
	}

	cookie := s.cookiesStore.CreateSession(user.ID)
	http.SetCookie(w, cookie)
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
