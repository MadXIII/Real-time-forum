package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"
)

// CreatePost - /newpost's handler
func (a *API) CreatePost(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	categories, err := a.service.GetCategories()
	// 	if err != nil {
	// 		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, GetCategories: %w", err))
	// 		return
	// 	}
	// 	bytes, err := json.Marshal(&categories)
	// 	if err != nil {
	// 		logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePostPage, Marshal: %w", err))
	// 	}
	// 	w.Write(bytes)
	// }

	if r.Method == http.MethodPost {
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

		if err = s.checkNewPostDatas(&newPost); err != nil {
			logger(w, http.StatusBadRequest, err)
			return
		}

		newPost.Username, err = s.getUsernameByCookie(r)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, getUsernameByCookie: %w", err))
			return
		}

		postID, err := s.store.InsertPost(&newPost)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleCreatePost, InsertPost: %w", err))
			return
		}

		// create object to Response about Success
		resp := struct {
			ID     int    `json:"id"`
			Notify string `json:"notify"`
		}{
			ID:     postID,
			Notify: "Post is Created",
		}

		bytes, err = json.Marshal(&resp)
		if err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}

		w.Write(bytes)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// handleCreatePost - if CreatePost POST method

// checkNewPostDatas - check Post Datas before Insert it into db
func (s *Server) checkNewPostDatas(post *model.Post) error {
	if err := s.store.CheckCategoryID(post.CategoryID); err != nil {
		return err
	}

	if len(post.Title) == 0 || len(post.Title) > 32 {
		return newErr.ErrPostTitle
	}

	if len(post.Content) == 0 {
		return newErr.ErrPostContent
	}
	// Set date format
	post.Timestamp = time.Now().Format("2.Jan.2006, 15:04")
	return nil
}

// getUsernameByCookie - get Username from db, by GetIDByCookie
func (s *Server) getUsernameByCookie(req *http.Request) (string, error) {
	ck, err := req.Cookie("session")
	if err != nil {
		return "", fmt.Errorf("getUsernameByCookie, r.Cookie(\"session\"): %w", err)
	}

	id, err := s.cookiesStore.GetIDByCookie(ck)
	if err != nil {
		return "", err
	}

	username, err := s.store.GetUsernameByID(id)
	if err != nil {
		return "", err
	}

	return username, nil
}
