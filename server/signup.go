package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

//SignUp page GET, POST
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.Parser()
		w.WriteHeader(http.StatusOK)
		if err = s.temp.Execute(w, nil); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
	} else if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		response := make(map[string]string)
		var newUser models.User
		var result string

		if err = json.Unmarshal(bytes, &newUser); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		result, ok := isEmpty(newUser)

		if ok {
			sendNotify(w, result, response, 400)
			return
		}

		if !isValidEmail(newUser.Email) {
			result = "Invalid Email"
			sendNotify(w, result, response, 400)
			return
		}

		if newUser.Password != newUser.Confirm {
			result = "Different second password"
			sendNotify(w, result, response, 400)
			return
		}

		if !isValidPass(newUser.Password) {
			result = "Invlaid Pass"
			sendNotify(w, result, response, 400)
			return
		}

		// hash, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }

		if err = s.store.InsertUser(newUser); err != nil {
			if strings.Contains(err.Error(), "nickname") {
				result = "Nickname is already in use"
				SendNotify(w, result, 500)
				return
			}
			if strings.Contains(err.Error(), "email") {
				result = "Email is already in use"
				SendNotify(w, result, 500)
				return
			}
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		//creatSession

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created succesfully"))
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
		return
	}
}

//send Notification to Front
func sendNotify(w http.ResponseWriter, result string, response map[string]string, status int) {
	response["notify"] = result
	notify, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	w.WriteHeader(status)
	w.Write(notify)
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
	var res string
	if newUser.Nickname == "" {
		res = "Nickname is empty"
		return res, true
	}
	if newUser.Email == "" {
		res = "Email is empty"
		return res, true
	}
	if newUser.Password == "" {
		res = "Password is empty"
		return res, true
	}
	if newUser.Confirm == "" {
		res = "Confirm is empty"
		return res, true
	}
	if newUser.FirstName == "" {
		res = "Firstname is empty"
		return res, true
	}
	if newUser.LastName == "" {
		res = "Lastname is empty"
		return res, true
	}
	if newUser.Gender == "" {
		res = "Gender is empty"
		return res, true
	}
	if newUser.Age == 0 {
		res = "Age is empty"
		return res, true
	}
	return res, false
}
