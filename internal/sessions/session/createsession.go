package session

import (
	newErr "forum/internal/error"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Store - to store cookies in map type
type Store struct {
	cookies map[int]*http.Cookie
	users   map[string]bool
}

// CreateSession - create session
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

// New - Initializtioner of cookie store
func NewSessionStore(usersList []string) *Store {
	s := new(Store)
	s.cookies = make(map[int]*http.Cookie)
	s.users = make(map[string]bool)
	for _, user := range usersList {
		s.users[user] = false
	}
	return s
}

// DeleteCoo`kie - delete cookie if find from map
func (s *Store) DeleteCookie(ck *http.Cookie) error {
	for key, val := range s.cookies {
		if val.Value == ck.Value {
			delete(s.cookies, key)
			return nil
		}
	}
	return newErr.ErrDelCookie
}

// CheckCookie - check cookie in map
func (s *Store) CheckCookie(cookieHash string) error {
	for _, r := range s.cookies {
		if r.Value == cookieHash {
			return nil
		}
	}
	return newErr.ErrNoCookie
}

// GetIDByCookie - search userid in cookies by request.Cookie
func (s *Store) GetIDByCookie(inpCookie *http.Cookie) (int, error) {
	for id, ck := range s.cookies {
		if ck.Value == inpCookie.Value {
			return id, nil
		}
	}
	return 0, newErr.ErrNoCookie
}

func (s *Store) SetOnlineUser(nickname string) {
	s.users[nickname] = true
}

func (s *Store) SetOfflineUser(nickname string) {
	s.users[nickname] = false
}

func (s *Store) GetOnlineList() map[string]bool {
	return s.users
}
