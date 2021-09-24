package session

import "net/http"

type Repository interface {
	CreateSession(int) *http.Cookie
}
