package session

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Store struct {
	cookies map[int]*http.Cookie
}

//CreateSession - create session
func (s *Store) CreateSession(uid int) *http.Cookie {
	sid := uuid.NewV4().String()
	s.cookies[uid] = &http.Cookie{
		Name:   "session",
		Value:  sid,
		Path:   "/",
		MaxAge: 86400,
	}
	return s.cookies[uid]
}
func New() *Store {
	s := new(Store)
	s.cookies = make(map[int]*http.Cookie)
	return s
}
