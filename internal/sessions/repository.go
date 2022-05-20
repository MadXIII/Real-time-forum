package session

import (
	"forum/internal/models"
	"net/http"
)

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
	AddOnlineUser(int, string)
	SetOnlineUser(string)
	SetOfflineUser(string)
	GetOnlineList() []models.OnlineUsers
}
