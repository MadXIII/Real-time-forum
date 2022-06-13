package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
)

type Post struct {
	repo repository.Repository
}

func NewPost(repo repository.Repository) *Post {
	return &Post{repo: repo}
}

func (p *Post) CheckData(post *model.Post) (int, int, error) {
	// if err := s.store.CheckCategoryID(post.CategoryID); err != nil {
	// 	return err
	// }

	if len(post.Title) == 0 || len(post.Title) > 32 {
		return 0, http.StatusBadRequest, newErr.ErrPostTitle
	}

	if len(post.Content) == 0 {
		return 0, http.StatusBadRequest, newErr.ErrPostContent
	}

	id, err := p.repo.CreatePost(post)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return id, 0, nil
}

// GetID - get PostID from request and check for digits
func (p *Post) Post(r *http.Request) (model.Post, error) {
	r.ParseForm()
	id := r.FormValue("id")

	var post model.Post

	_, err := strconv.Atoi(id)
	if err != nil {
		return post, fmt.Errorf("GetID, Atoi: %w", err)
	}

	post, err = p.repo.GetPostByID(id)
	if err != nil {
		return post, err
	}
	return post, nil
}

// checkInsertComment - get Username from request, check length of comment and Insert it into db if correct
func (p *Post) CheckComment(comment *model.Comment) (status int, err error) {
	if len(comment.Content) < 1 {
		return http.StatusBadRequest, newErr.ErrEmptyComment
	}
	if len(comment.Content) > 256 {
		return http.StatusBadRequest, newErr.ErrLenComment
	}

	// set date to comment
	comment.Timestamp = time.Now().Format("2.Jan.2006, 15:04")

	if err = p.repo.InsertComment(comment); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, err
}

func (p *Post) CheckVote(like *model.PostLike) (status int, err error) {
	like.VoteState, err = p.repo.VoteState(like)
	if err != nil {
		if err = p.repo.CreateVote(like); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	if err = p.voteThumbler(like); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = p.repo.UpdateVotes(like); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (p *Post) voteThumbler(like *model.PostLike) (err error) {
	if like.VoteState {
		like.VoteState = false
		if err = p.repo.UpdateVoteState(like); err != nil {
			return err
		}
	} else {
		like.VoteState = true
		if err = p.repo.UpdateVoteState(like); err != nil {
			return err
		}
	}
	return nil
}

func (p *Post) Comments(id int) ([]model.Comment, error) {
	comments, err := p.repo.PostComments(id)
	return comments, err
}
