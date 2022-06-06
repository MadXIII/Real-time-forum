package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"unicode"

	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.Repository
}

func NewUser(repo repository.Repository) *User {
	return &User{repo: repo}
}

func (u *User) CheckUserData(user model.User) error {
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

	bytes, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreateAccount, GenerateFromPassword: %w", err))
		return
	}
	user.Password = string(bytes)

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
func checkEmpty(newUser model.User) error {
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
