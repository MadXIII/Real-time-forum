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

	fmt.Println(post)

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
		return
	}

	data := struct {
		models.Comment
		models.PostLike
		VoteType string `json:"voteType"`
	}{}

	if err = json.Unmarshal(bytes, &data); err != nil {
		logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, Unmarshal(newPost): %w", err))
		return
	}
	//create on func for this
	if data.Comment.PostID != 0 {
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
	}
	if data.VoteType == "like" {
		if err = s.likeThumbler(r, &data.PostLike); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		s.store.ChangeLikeDislikeDiff(data.PostLike.PostID, data.PostLike.VoteState)
	}
	if data.VoteType == "dislike" {
		if err = s.likeThumbler(r, &data.PostLike); err != nil {
			logger(w, http.StatusInternalServerError, err)
			return
		}
		s.store.ChangeLikeDislikeDiff(data.PostLike.PostID, data.PostLike.VoteState)
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

func (s *Server) likeThumbler(req *http.Request, like *models.PostLike) (err error) { // ERROR :=
	like.UserID, err = s.cookiesStore.GetIDByCookie(req)
	if err != nil {
		return err
	}
	like.VoteState, err = s.store.GetVoteState(like.PostID, like.UserID)
	if err != nil {
		like.VoteState = true
		if err = s.store.InsertLike(like); err != nil {
			return err
		}
		return nil
	}
	if like.VoteState == true {
		like.VoteState = false
		s.store.UpdateVoteState(like)

	} else {
		like.VoteState = true
		s.store.UpdateVoteState(like)

	}
	return nil
}
