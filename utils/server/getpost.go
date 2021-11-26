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
		Post     models.Post       `json:"Post"`
		Comments []models.Comment  `json:"Comments"`
		Likes    []models.PostLike `json:"Like"`
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
		return
	}

	data := struct {
		models.Comment
		models.PostLike
	}{}

	if err = json.Unmarshal(bytes, &data); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, Unmarshal(newPost): %w", err))
		return
	}

	if err = checkComment(data.Comment.Content); err != nil {
		logger(w, http.StatusBadRequest, err)
		return
	}

	//Set date format
	data.Comment.Timestamp = time.Now().Format("2.Jan.2006, 15:04")

	data.Comment.Username, err = s.getUsernameByCookie(r)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, getUsernameByCookie: %w", err))
		return
	}

	if err := s.store.InsertComment(&data.Comment); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, InsertComment: %w", err))
		return
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

// func (s *Server) checkLike(r *http.Request, like *models.PostLike) (err error) {
// 	like.PostID, err = s.cookiesStore.GetIDByCookie(r)
// 	if err != nil {
// 		return err
// 	}
// 	if like.VoteState != 1 {

// 	}
// 	return nil
// }
