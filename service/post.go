package service

import (
	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
)

type Post struct {
	repo repository.Repository
}

func NewPost(repo repository.Repository) *Post {
	return &Post{repo: repo}
}

func (p *Post) Create(post *model.Post) (int, int, error) {
	return 0, 0, nil
}
