package session

import "net/http"

//Repository - interface to work with cookies
type Repository interface {
	CreateSession(int) *http.Cookie
	DeleteCookie(*http.Cookie) error
	CheckCookie(string) error
	GetCookies() map[int]*http.Cookie
	GetIdByCookie(*http.Request) int
	// LogOut(http.ResponseWriter, *http.Request)
}
