package server

import (
	"encoding/json"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

//SignUp page GET, POST
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handleCreateAccount(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

//handleCreateAccount - if SignUp POST method
func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
	var newUser models.User

	if err = json.Unmarshal(bytes, &newUser); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}

	if err := isCorrectDatasToSignUp(newUser); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	bytes, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
	newUser.Password = string(bytes)

	if err = s.insertUserDB(newUser); err != nil {
		logger(w, http.StatusBadRequest, err)
	}

	cookie := s.cookiesStore.CreateSession(newUser.ID)

	http.SetCookie(w, cookie)
}

//insertUserDB - Insert User in DB if no error
func (s *Server) insertUserDB(user models.User) error {
	if err := s.store.InsertUser(user); err != nil {
		if strings.Contains(err.Error(), "nickname") {
			return newErr.ErrNickname
		}
		if strings.Contains(err.Error(), "email") {
			return newErr.ErrEmail
		}
		return err
	}
	return nil
}

//isCorrcetDataToSignUp ...
func isCorrectDatasToSignUp(user models.User) error {
	if err := checkEmpty(user); err != nil {
		return err
	}
	if !isValidEmail(user.Email) {
		return newErr.ErrInvalidEmail
	}
	if user.Password != user.Confirm {
		return newErr.ErrDiffSecondPass
	}
	if !isValidPass(user.Password) {
		return newErr.ErrInvalidPass
	}
	return nil
}

//checkin email for validity
func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_.]+[^\^!#\$%&'\@()*+\/=\?\^\n_{\|}~-]@[a-z]{2,}\.[a-zA-Z]{2,6}$`)
	if !regex.MatchString(email) {
		return false
	}
	return true
}

//checking pass for validity
func isValidPass(pass string) bool {
	if len(pass) < 8 || len(pass) > 32 {
		return false
	}
	for _, r := range pass {
		if r < 33 || r > 126 {
			return false
		}
	}
	var low, up, num bool
	for _, r := range pass {
		if unicode.IsLower(r) {
			low = true
		}
		if unicode.IsUpper(r) {
			up = true
		}
		if unicode.IsNumber(r) {
			num = true
		}
	}
	return low && up && num
}

//checking for empty fields in signup page
func checkEmpty(newUser models.User) error {
	if newUser.Nickname == "" {
		return newErr.ErrEmptyNickname
	}
	if newUser.Email == "" {
		return newErr.ErrEmptyEmail
	}
	if newUser.Password == "" {
		return newErr.ErrEmptyPassword
	}
	if newUser.Confirm == "" {
		return newErr.ErrEmptyConfirm
	}
	if newUser.FirstName == "" {
		return newErr.ErrEmptyFirstname
	}
	if newUser.LastName == "" {
		return newErr.ErrEmptyLastname
	}
	if newUser.Gender == "" {
		return newErr.ErrEmptyGender
	}
	if newUser.Age == 0 {
		return newErr.ErrEmptyAge
	}
	return nil
}
