package server

import (
	"encoding/json"
	"fmt"
	"forum/internal/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	newErr "forum/internal/error"
)

//GetPost - /post?id=... handler
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

//handleGetPostPage - if GetPost GET method
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

	bytes, err := json.Marshal(&PostPageData)
	if err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handleGetPostPage, Marshal(PostData): %w", err))
		return
	}

	w.Write(bytes)
	return
}

//handlePost - if GetPost POST method
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
		logger(w, http.StatusBadRequest, fmt.Errorf("handlePost, Unmarshal(newPost): %w", err))
		return
	}

	if data.Comment.PostID != 0 {
		status, err := s.checkInsertComment(r, &data.Comment)
		if err != nil {
			logger(w, status, err)
			return
		}
	}

	if data.PostLike.PostID != 0 {
		if status, err := s.checkInsertUpdVote(r, &data.PostLike); err != nil {
			logger(w, status, err)
			return
		}
	}
	success(w, "Comment is created")
}

//checkAndGetPostID - get PostID from request and check for digits
func checkAndGetPostID(r *http.Request) (string, error) {
	r.ParseForm()
	value := r.FormValue("id")

	_, err := strconv.Atoi(value)
	if err != nil {
		return "", fmt.Errorf("checkAndGetPostID, Atoi: %w", err)
	}
	return value, nil
}

//checkInsertComment - get Username from request, check length of comment and Insert it into db if correct
func (s *Server) checkInsertComment(req *http.Request, comment *models.Comment) (status int, err error) {
	comment.Username, err = s.getUsernameByCookie(req)
	if err != nil {
		return http.StatusUnauthorized, newErr.ErrUnsignComment
	}

	if len(comment.Content) < 1 {
		return http.StatusBadRequest, newErr.ErrEmptyComment
	}
	if len(comment.Content) > 256 {
		return http.StatusBadRequest, newErr.ErrLenComment
	}

	//set date to comment
	comment.Timestamp = time.Now().Format("2.Jan.2006, 15:04")

	if err = s.store.InsertComment(comment); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, err
}

//checkInsertUpdVote - get UserID from request, Insert it into db if it first request, else set vote state
func (s *Server) checkInsertUpdVote(req *http.Request, like *models.PostLike) (status int, err error) {
	ck, err := req.Cookie("session")
	if err != nil {
		return http.StatusUnauthorized, newErr.ErrUnsignVote
	}
	like.UserID, err = s.cookiesStore.GetIDByCookie(ck)
	if err != nil || like.UserID < 1 {
		return http.StatusUnauthorized, newErr.ErrUnsignVote
	}

	like.VoteState, err = s.store.GetVoteState(like)
	if err != nil {
		if err = s.store.InsertVote(like); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	if err = s.voteThumbler(like); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = s.store.UpdateVotes(like); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

//voteThumbler - thumbler for vote state
func (s *Server) voteThumbler(like *models.PostLike) (err error) {
	if like.VoteState == true {
		like.VoteState = false
		if err = s.store.UpdateVoteState(like); err != nil {
			return err
		}
	} else {
		like.VoteState = true
		if err = s.store.UpdateVoteState(like); err != nil {
			return err
		}
	}
	return nil
}
