package session

import "net/http"

// Repository - interface to work with cookies
type Repository interface {
	Cookies
	UserList
}

type Cookies interface {
	CreateSession(int) *http.Cookie
	DeleteCookie(*http.Cookie) error
	CheckCookie(string) error
	GetIDByCookie(*http.Cookie) (int, error)
}

type UserList interface {
	SetOnlineUser(string)
	SetOfflineUser(string)
	GetOnlineList() map[string]bool
}
