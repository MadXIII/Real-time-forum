package testsession

import "net/http"

type TestSession struct {
}

func (s *TestSession) CreateSession(int) *http.Cookie {
	return &http.Cookie{}
}
func (s *TestSession) DeleteCookie(*http.Cookie) error {
	return nil
}
func (s *TestSession) CheckCookie(string) error {
	return nil
}
func (s *TestSession) GetIDByCookie(*http.Request) int {
	return 0
}
