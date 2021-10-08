package session

import "net/http"

//Repository - interface to work with cookies
type Repository interface {
	CreateSession(int) *http.Cookie
	DeleteCookie(*http.Cookie) error
	CheckCookie(string) error
	// GetCookies() map[int]*http.Cookie
	// LogOut(http.ResponseWriter, *http.Request)
}
