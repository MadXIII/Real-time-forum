package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Sign struct {
	NickOrEmail string `json:"nickoremail"`
	Password    string `json:"password"`
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.Parser()
		if err = s.temp.Execute(w, nil); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
	} else if r.Method == http.MethodPost {
		var signer Sign
		var user models.User
		var result string

		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if err = json.Unmarshal(bytes, &signer); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if strings.Contains(signer.NickOrEmail, "@") {
			user, err = s.store.GetUserByEmail(signer.NickOrEmail)
			if err != nil {
				result = "Wrong Nickname or Email"
				SendNotify(w, result, 400)
				log.Println(err)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password))
			if err != nil {
				result = "Wrong Password"
				SendNotify(w, result, 400)
				log.Println(err)
				return
			}

		} else {
			user, err = s.store.GetUserByNickname(signer.NickOrEmail)
			if err != nil {
				result = "Wrong Nickname or Email"
				SendNotify(w, result, 400)
				log.Println(err)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password))
			if err != nil {
				result = "Wrong Password"
				SendNotify(w, result, 400)
				log.Println(err)
				return
			}
		}
		CreateSession(w, user.ID)

		w.WriteHeader(200)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
		return
	}
}
