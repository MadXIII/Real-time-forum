package session

import (
	newErr "forum/internal/error"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//Store - to store cookies in map type
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

//New - Initializtioner of cookie store
func New() *Store {
	s := new(Store)
	s.cookies = make(map[int]*http.Cookie)
	return s
}

//DeleteCoo`kie - delete cookie if find from map
func (s *Store) DeleteCookie(ck *http.Cookie) error {
	for key, val := range s.cookies {
		if val.Value == ck.Value {
			delete(s.cookies, key)
			return nil
		}
	}
	return newErr.ErrDelCookie
}

//CheckCookie - check cookie in map
func (s *Store) CheckCookie(cookieHash string) error {
	for _, r := range s.cookies {
		if r.Value == cookieHash {
			return nil
		}
	}
	return newErr.ErrNoCookie
}

//GetIDByCookie - search userid in cookies by request.Cookie
func (s *Store) GetIDByCookie(inpCookie *http.Cookie) (int, error) {
	for id, ck := range s.cookies {
		if ck.Value == inpCookie.Value {
			return id, nil
		}
	}
	return 0, newErr.ErrNoCookie
}
