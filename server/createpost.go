package server

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"io/ioutil"
	"net/http"
	"time"
)

//CreatePost ...
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp := Parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		var newPost models.Post
		newPost.Timestamp = time.Now()

		if err = json.Unmarshal(bytes, &newPost); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		fmt.Println(newPost)
		if err = s.store.InsertPost(newPost); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		// w.WriteHeader(200)

	}
}
