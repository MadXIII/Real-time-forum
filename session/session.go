package session

import (
	"net/http"
)

type Sesssion struct {
	Cookie
	// Chat
}

type Cookie interface {
	CreateSession(int) *http.Cookie
	DeleteCookie(*http.Cookie) error
	CheckCookie(string) error
	GetIDByCookie(*http.Cookie) (int, error)
}

// type Chat interface {
// 	AddOnlineUser(int, string)
// 	SetOnlineUser(string)
// 	SetOfflineUser(string)
// 	GetOnlineList() []model.OnlineUsers
// }

func New() *Sesssion {
	return &Sesssion{
		Cookie: NewCookie(),
		// Chat: NewChat(),
	}
}
