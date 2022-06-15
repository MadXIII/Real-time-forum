package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"
)

// CreatePost - /newpost's handler
func (a *API) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categories, err := a.service.Categories()
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, GetCategories: %w", err))
			return
		}
		bytes, err := json.Marshal(&categories)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, Marshal: %w", err))
		}
		w.Write(bytes)

	case http.MethodPost:
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, ReadAll(r.Body): %w", err))
			return
		}

		if len(bytes) == 0 {
			logger(w, http.StatusBadRequest, newErr.ErrNilBody)
			return
		}
		var newPost model.Post
		if err = json.Unmarshal(bytes, &newPost); err != nil {
			logger(w, http.StatusBadRequest, fmt.Errorf("handleCreatePost, Unmarshal(newPost): %w", err))
			return
		}

		newPost.Username, err = a.service.GetUsernameByCookie(r)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		id, status, err := a.service.CheckData(&newPost)
		if err != nil {
			logger(w, status, err)
		}

		// create object to Response about Success
		resp := struct {
			ID     int    `json:"id"`
			Notify string `json:"notify"`
		}{
			ID:     id,
			Notify: "Post is Created",
		}

		bytes, err = json.Marshal(&resp)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		w.Write(bytes)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
