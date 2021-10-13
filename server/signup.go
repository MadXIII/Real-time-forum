package server

import (
	"encoding/json"
	"forum/models"
	"forum/sessions/session"
	"io/ioutil"
	"log"
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
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		var newUser models.User

		if err = json.Unmarshal(bytes, &newUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if res, ok := isCorrcetDatasToSignUp(newUser); !ok {
			sendNotify(w, res, http.StatusBadRequest)
			return
		}

		bytes, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		newUser.Password = string(bytes)

		if err = s.store.InsertUser(newUser); err != nil {
			if strings.Contains(err.Error(), "nickname") {
				sendNotify(w, "Nickname is already in use", http.StatusInternalServerError)
				return
			}
			if strings.Contains(err.Error(), "email") {
				sendNotify(w, "Email is already in use", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		cookie := session.New().CreateSession(newUser.ID)

		http.SetCookie(w, cookie)
		// log.Println(cookie)

		//connect session from private browser

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created succesfully"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
		return
	}
}

//send Notification to Front
func sendNotify(w http.ResponseWriter, result string, status int) {
	response := make(map[string]string)
	response["notify"] = result
	notify, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(status)
	w.Write(notify)
}

//isCorrcetDataToSignUp ...
func isCorrcetDatasToSignUp(user models.User) (string, bool) {
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
