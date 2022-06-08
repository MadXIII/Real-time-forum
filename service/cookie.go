package service

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	newErr "github.com/madxiii/real-time-forum/error"
)

type Store struct {
	cookie map[int]*http.Cookie
}

func NewStore() *Store {
	return &Store{
		cookie: map[int]*http.Cookie{},
	}
}

func (s *Store) Create(id int) (*http.Cookie, error) {
	sid, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("CreateSession, uuid.NewV4: %s", err.Error())
	}
	s.cookie[id] = &http.Cookie{
		Name:   "session",
		Value:  sid.String(),
		Path:   "/",
		MaxAge: 86400,
	}
	return s.cookie[id], nil
}

func (s *Store) Delete(ck *http.Cookie) error {
	for key, val := range s.cookie {
		if val.Value == ck.Value {
			delete(s.cookie, key)
			return nil
		}
	}

	return newErr.ErrDelCookie
}

// Check - check cookie in map
func (s *Store) Check(cookieHash string) error {
	for _, r := range s.cookie {
		if r.Value == cookieHash {
			return nil
		}
	}
	return newErr.ErrNoCookie
}

// GetIDByCookie - search userid in cookie by request.Cookie
func (s *Store) GetIDByCookie(inpCookie *http.Cookie) (int, error) {
	for id, ck := range s.cookie {
		if ck.Value == inpCookie.Value {
			return id, nil
		}
	}
	return 0, newErr.ErrNoCookie
}
