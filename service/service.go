package service

import (
	"net/http"

	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
)

type Service struct {
	Registerer
	Loger
	Cookie
	PostChecker
}

type Registerer interface {
	Register(user *model.User) (int, error)
}

type Loger interface {
	Login(user model.Sign) (int, int, error)
	Logout(http.ResponseWriter, *http.Cookie)
}

type Cookie interface {
	Create(id int) (*http.Cookie, error)
	Delete(*http.Cookie) error
	Check(string) error
	GetIDByCookie(*http.Cookie) (int, error)
	GetUsernameByCookie(*http.Request) (string, error)
}

type PostChecker interface {
	CheckData(*model.Post) (int, int, error)
	Post(*http.Request) (model.Post, error)
	CheckComment(*model.Comment) (int, error)
	CheckVote(*model.PostLike) (int, error)
	Comments(int) ([]model.Comment, error)
	Categories() ([]model.Categories, error)
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Registerer:  NewUser(*repo),
		Loger:       NewLog(*repo),
		Cookie:      NewStore(*repo),
		PostChecker: NewPost(*repo),
	}
}
