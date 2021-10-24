package server

import (
	"encoding/json"
	newErr "forum/internal/error"
	"forum/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

//SignUp page GET, POST
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	if r.Method == http.MethodPost {
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
			SendNotify(w, err.Error(), http.StatusBadRequest)
			return
		}

		bytes, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		newUser.Password = string(bytes)

		//move outside body
		if err = s.store.InsertUser(newUser); err != nil {
			if strings.Contains(err.Error(), "nickname") {
				SendNotify(w, "Nickname is already in use", http.StatusInternalServerError)
				return
			}
			if strings.Contains(err.Error(), "email") {
				SendNotify(w, "Email is already in use", http.StatusInternalServerError)
				return
			}
			logger(w, http.StatusInternalServerError, err)
			return
		}

		cookie := s.cookiesStore.CreateSession(newUser.ID)

		http.SetCookie(w, cookie)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created succesfully"))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 method not allowed"))
	return
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
