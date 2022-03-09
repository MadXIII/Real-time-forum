package session

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository - interface to work with cookies
type Repository interface {
	CreateSession(primitive.ObjectID) *http.Cookie
	DeleteCookie(*http.Cookie) error
	CheckCookie(string) error
	GetIDByCookie(*http.Cookie) (int, error)
}
