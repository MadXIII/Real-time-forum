package session

import (
	"fmt"
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

//DeleteCookie - delete cookie if find from map
func (s *Store) DeleteCookie(ck *http.Cookie) error {
	for key, val := range s.cookies {
		if val.Value == ck.Value {
			delete(s.cookies, key)
			return nil
		}
	}
	return fmt.Errorf("Something wrong with cookies to delete it")
}

//CheckCookie - check cookie in map
func (s *Store) CheckCookie(cookieHash string) error {
	for _, r := range s.cookies {
		if r.Value == cookieHash {
			return nil
		}
	}
	return fmt.Errorf("Problem with cookie")
}

//GetIdByCookie - search userid in cookies by request.Cookie
func (s *Store) GetIdByCookie(req *http.Request) int {
	userCk, err := req.Cookie("session")
	if err != nil {
		return -1
	}
	for id, ck := range s.cookies {
		if ck.Value == userCk.Value {
			return id
		}
	}
	return -1
}
