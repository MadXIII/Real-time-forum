package service

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/repository"
)

type Store struct {
	repo   repository.Repository
	cookie map[int]*http.Cookie
}

func NewStore(repo repository.Repository) *Store {
	return &Store{
		repo:   repo,
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

// getUsernameByCookie - get Username from db, by GetIDByCookie
func (s *Store) GetUsernameByCookie(req *http.Request) (string, error) {
	ck, err := req.Cookie("session")
	if err != nil {
		return "", fmt.Errorf("GetUsernameByCookie, r.Cookie: %w", err)
	}

	id, err := s.GetIDByCookie(ck)
	if err != nil {
		return "", err
	}

	username, err := s.repo.GetUsernameByID(id)
	if err != nil {
		return "", err
	}

	return username, nil
}
