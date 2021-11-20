package server

import (
	"encoding/json"
	"fmt"
	"forum/utils/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	newErr "forum/utils/internal/error"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handleGetPostPage(w, r)
		return
	}
	if r.Method == http.MethodPost {
		s.handlePost(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (s *Server) handleGetPostPage(w http.ResponseWriter, r *http.Request) {
	postid, err := checkAndGetPostID(r)
	if err != nil {
		logger(w, http.StatusBadRequest, fmt.Errorf("handleGetPostPage, checkAndGetPostID: %w", err))
		return
	}

	post, err := s.store.GetPostByID(postid)
	if err != nil {
		logger(w, http.StatusBadRequest, fmt.Errorf("handleGetPostPage, GetPostByID: %w", err))
		return
	}

	comments, err := s.store.GetCommentsByPostID(postid)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleGetPostPage, GetCommentsByPostID: %w", err))
		return
	}

	PostPageData := struct {
		Post     models.Post      `json:"Post"`
		Comments []models.Comment `json:"Comments"`
	}{
		Post:     post,
		Comments: comments,
	}

	bytes, err := json.Marshal(PostPageData)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleGetPostPage, Marshal(PostData): %w", err))
		return
	}

	w.Write(bytes)
	return
}

//what about postID? Still foreign Key or just get from URL.path?

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, ReadAll(r.Body): %w", err))
	}

	var newComment models.Comment

	if err = json.Unmarshal(bytes, &newComment); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, Unmarshal(newPost): %w", err))
	}

	if err = checkComment(newComment.Content); err != nil {
		logger(w, http.StatusBadRequest, err)
	}

	//Set date format
	newComment.Timestamp = time.Now().Format("2.Jan.2006, 15:04")

	newComment.Username, err = s.getUsernameByCookie(r)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, getUsernameByCookie: %w", err))
	}

	if err := s.store.InsertComment(&newComment); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, InsertComment: %w", err))
	}

	success(w, "Comment is created")
}

func checkAndGetPostID(r *http.Request) (string, error) {
	r.ParseForm()
	value := r.FormValue("id")

	_, err := strconv.Atoi(value)
	if err != nil {
		return "", fmt.Errorf("checkAndGetPostID, Atoi: %w", err)
	}
	return value, nil
}

func checkComment(comment string) error {
	if len(comment) < 1 {
		return newErr.ErrEmptyComment
	}
	if len(comment) > 256 {
		return newErr.ErrLenComment
	}
	return nil
}
