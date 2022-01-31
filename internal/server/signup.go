package server

import (
	"encoding/json"
	"fmt"
	newErr "forum/internal/error"
	"forum/internal/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreateAccount, ReadAll(r.Body): %w", err))
		return
	}
	var newUser models.User

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
}

//insertUserDB - Insert User in DB if no error
func (s *Server) insertUserDB(user *models.User) error {
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
	if err := isValidEmail(user.Email); err != nil {
		return err
	}
	if user.Password != user.Confirm {
		return newErr.ErrDiffSecondPass
	}
	if err := isValidPass(user.Password); err != nil {
		return err
	}
	if err := isValidAge(user.Age); err != nil {
		return err
	}
	return nil
}

//checkin email for validity
func isValidEmail(email string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_.]+[^\^!#\$%&'\@()*+\/=\?\^\n_{\|}~-]@[a-z]{2,}\.[a-zA-Z]{2,6}$`)
	if !regex.MatchString(email) {
		return newErr.ErrInvalidEmail
	}
	return nil
}

//checking pass for validity
func isValidPass(pass string) error {
	if len(pass) < 8 || len(pass) > 32 {
		return newErr.ErrInvalidPass
	}
	for _, r := range pass {
		if r < 33 || r > 126 {
			return newErr.ErrInvalidPass
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
	if low && up && num {
		return nil
	}
	return newErr.ErrInvalidPass
}

//isValidAge - is input age is valid?
func isValidAge(age string) error {
	if age == "" {
		return nil
	}
	digit, err := strconv.Atoi(age)
	if err != nil {
		return newErr.ErrInvalidAge
	}
	if digit < 6 || digit > 100 {
		return newErr.ErrInvalidAge
	}
	return nil
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
	return nil
}
