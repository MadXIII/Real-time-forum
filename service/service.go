package service

import (
	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository"
)

type Service struct {
	Register
}

type Register interface {
	CheckUserData(user model.User) error
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Register: NewUser(*repo),
	}
}
