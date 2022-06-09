package service

import (
	"net/http"

	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
)

type Service struct {
	Registerer
	Loginer
	Logouter
	Cookie
	Post
}

type Registerer interface {
	Register(user *model.User) (int, error)
}

type Loginer interface {
	Login(user model.Sign) (int, int, error)
}

type Logouter interface {
	Logout(http.ResponseWriter, *http.Cookie)
}

type Cookie interface {
	Create(id int) (*http.Cookie, error)
	Delete(*http.Cookie) error
	Check(string) error
	GetIDByCookie(*http.Cookie) (int, error)
}

type Post interface {
	Create(*model.Post) (int, int, error)
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Registerer: NewUser(*repo),
		Loginer:    NewLogin(*repo),
		Logouter:   NewLogout(),
		Cookie:     NewStore(),
		Post:       NewPost(repo),
	}
}
