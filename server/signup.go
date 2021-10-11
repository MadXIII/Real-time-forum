package server

import (
	"encoding/json"
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

		if res, ok := isCorrectDatasToSignUp(newUser); !ok {
			SendNotify(w, res, http.StatusBadRequest)
			return
		}

		bytes, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		newUser.Password = string(bytes)

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
func isCorrectDatasToSignUp(user models.User) (string, bool) {
	if res, ok := isEmpty(user); ok {
		return res, false
	}
	if !isValidEmail(user.Email) {
		return "Invalid Email", false
	}
	if user.Password != user.Confirm {
		return "Different second password", false
	}
	if !isValidPass(user.Password) {
		return "Invlaid Pass", false
	}
	return "", true
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
	var low, up, num bool
	if len(pass) < 8 || len(pass) > 32 {
		return false
	}
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

//checking for emptys in signup page
func isEmpty(newUser models.User) (string, bool) {
	if newUser.Nickname == "" {
		return "Nickname is empty", true
	}
	if newUser.Email == "" {
		return "Email is empty", true
	}
	if newUser.Password == "" {
		return "Password is empty", true
	}
	if newUser.Confirm == "" {
		return "Confirm is empty", true
	}
	if newUser.FirstName == "" {
		return "Firstname is empty", true
	}
	if newUser.LastName == "" {
		return "Lastname is empty", true
	}
	if newUser.Gender == "" {
		return "Gender is empty", true
	}
	if newUser.Age == 0 {
		return "Age is empty", true
	}
	return "", false
}
