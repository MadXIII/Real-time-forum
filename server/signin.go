package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//Sign - struct to get Signer datas
type Sign struct {
	NickOrEmail string `json:"nickoremail"`
	Password    string `json:"password"`
}

//SignIn - Sigin page
func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}
	if r.Method == http.MethodPost {
		var signer Sign
		var user *models.User

		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err = json.Unmarshal(bytes, &signer); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		user, err = s.store.GetUserByLogin(signer.NickOrEmail)
		if err != nil {
			SendNotify(w, "Wrong Nickname or Email", http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signer.Password))
		if err != nil {
			SendNotify(w, "Wrong Password", http.StatusBadRequest)
			log.Println(err)
			return
		}

		cookie := s.cookiesStore.CreateSession(user.ID)

		http.SetCookie(w, cookie)

		w.WriteHeader(200)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 method not allowed"))
	return
}
