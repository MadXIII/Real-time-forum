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

	bytes, err := json.Marshal(&PostPageData)
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
	//create func for this
	if data.Comment.PostID != 0 {
		status, err := s.checkInsertComment(r, &data.Comment)
		if err != nil {
			logger(w, status, err)
			return
		}
	}

	if data.PostLike.PostID != 0 {
		if err = s.checkInsertUpdVote(r, &data.PostLike); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
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

func (s *Server) checkInsertComment(req *http.Request, comment *models.Comment) (status int, err error) {
	comment.Username, err = s.getUsernameByCookie(req)
	if err != nil {
		return http.StatusInternalServerError, newErr.ErrUnsignComment
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

func (s *Server) checkInsertUpdVote(req *http.Request, like *models.PostLike) (err error) {
	like.UserID, err = s.cookiesStore.GetIDByCookie(req)
	if err != nil || like.UserID < 1 {
		return newErr.ErrUnsignVote
	}

	like.VoteState, err = s.store.GetVoteState(like)
	if err != nil {
		if err = s.store.InsertLike(like); err != nil {
			return err
		}
	}

	if err = s.voteThumbler(like); err != nil {
		return err
	}

	s.store.UpdateLikes(like)

	return nil
}

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
