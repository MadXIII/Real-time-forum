package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"
)

func (a *API) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		post, err := a.service.Post(r)
		if err != nil {
			logger(w, http.StatusBadRequest, fmt.Errorf("Post, GetID: %w", err))
			return
		}

		// comments, err := s.store.GetCommentsByPostID(postid)
		// if err != nil {
		// 	logger(w, http.StatusInternalServerError, fmt.Errorf("handleGetPostPage, GetCommentsByPostID: %w", err))
		// 	return
		// }

		PostData := struct {
			Post     model.Post      `json:"Post"`
			Comments []model.Comment `json:"Comments"`
		}{
			Post: post,
			// Comments: comments,
		}

		bytes, err := json.Marshal(&PostData)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handleGetPostPage, Marshal(PostData): %w", err))
			return
		}

		w.Write(bytes)
		return
	}
	if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("handlePost, ReadAll(r.Body): %w", err))
			return
		}

		data := struct {
			model.Comment
			model.PostLike
		}{}

		if err = json.Unmarshal(bytes, &data); err != nil {
			logger(w, http.StatusBadRequest, fmt.Errorf("handlePost, Unmarshal(newPost): %w", err))
			return
		}

		username, err := a.service.GetUsernameByCookie(r)
		if err != nil {
			logger(w, http.StatusUnauthorized, newErr.ErrUnsignAction)
			return
		}
		ck, err := r.Cookie("session")
		if err != nil {
			logger(w, http.StatusUnauthorized, newErr.ErrUnsignAction)
			return
		}
		data.PostLike.UserID, err = a.service.GetIDByCookie(ck)
		if err != nil {
			logger(w, http.StatusUnauthorized, newErr.ErrUnsignAction)
			return
		}
		if data.Comment.PostID != 0 {
			data.Comment.Username = username
			status, err := a.service.CheckComment(&data.Comment)
			if err != nil {
				logger(w, status, err)
				return
			}
		} else {
			if status, err := a.service.CheckVote(&data.PostLike); err != nil {
				logger(w, status, err)
				return
			}
		}
		success(w, "Comment is created")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
