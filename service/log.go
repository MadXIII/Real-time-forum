package service

import (
	"net/http"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
	"golang.org/x/crypto/bcrypt"
)

type Log struct {
	repo repository.Repository
}

func NewLog(repo repository.Repository) *Log {
	return &Log{repo: repo}
}

func (l *Log) Login(signer model.Sign) (int, int, error) {
	if err := checkLoginDatas(signer); err != nil {
		return 0, http.StatusBadRequest, err
	}

	user, err := l.repo.User.GetUserByLogin(signer.Login)
	if err != nil {
		return 0, http.StatusBadRequest, newErr.ErrWrongLogin
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password)); err != nil {
		return 0, http.StatusBadRequest, newErr.ErrWrongPass
	}

	return user.ID, 0, nil
}

// checkLoginDatas - is empty or too long login datas
func checkLoginDatas(user model.Sign) error {
	if len(user.Login) == 0 || len(user.Login) > 32 {
		return newErr.ErrLoginData
	}
	if len(user.Password) < 8 || len(user.Password) > 32 {
		return newErr.ErrPassData
	}
	return nil
}

func (l *Log) Logout(w http.ResponseWriter, ck *http.Cookie) {
	ck.MaxAge = -1
	http.SetCookie(w, ck)
	w.WriteHeader(http.StatusOK)
}
