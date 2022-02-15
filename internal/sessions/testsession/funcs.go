package testsession

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type TestSession struct {
	mock.Mock
}

func (s *TestSession) CreateSession(uid int) *http.Cookie {
	args := s.Called(uid)
	return args.Get(0).(*http.Cookie)
}
func (s *TestSession) DeleteCookie(ck *http.Cookie) error {
	args := s.Called(ck)
	return args.Error(0)
}
func (s *TestSession) CheckCookie(hash string) error {
	args := s.Called(hash)
	return args.Error(0)
}
func (s *TestSession) GetIDByCookie(req *http.Request) (int, error) {
	args := s.Called(req)
	return args.Int(0), args.Error(1)
}
