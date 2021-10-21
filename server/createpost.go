package server

import (
	"encoding/json"
	"fmt"
	newErr "forum/internal/error"
	"forum/models"
	"io/ioutil"
	"log"
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
		log.Println(r.Body)
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		if len(bytes) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Request Body = nil")
			return
		}
		var newPost models.Post
		newPost.Timestamp = time.Now()

		if err = json.Unmarshal(bytes, &newPost); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		if err = checkNewPostDatas(newPost); err != nil {
			SendNotify(w, err.Error(), http.StatusBadRequest)
			return
		}

		//cant find how to error
		if newPost.UserID = s.cookiesStore.GetIdByCookie(r); newPost.UserID < 0 {
			logger(w, http.StatusInternalServerError, newErr.ErrNoCookie)
			return
		}
		//cant find how to error
		if err = s.store.InsertPost(newPost); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		fmt.Println(newPost)
		w.WriteHeader(200)
	}
}

func checkNewPostDatas(post models.Post) error {
	if len(post.Title) == 0 || len(post.Title) > 32 {
		return newErr.ErrPostTitle
	}
	if len(post.Content) == 0 {
		return newErr.ErrPostContent
	}
	return nil
}
