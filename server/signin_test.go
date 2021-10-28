package server

import (
	"bytes"
	"forum/database/testdb"
	"forum/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {
	mysrv := Init(&testdb.TestDB{}, &testsession.TestSession{})
	mysrv.router.HandleFunc("/signin", mysrv.SignIn)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method     string
		inputBody  []byte
		wantStatus int
	}{
		"Succes GET": {
			method:     "GET",
			inputBody:  nil,
			wantStatus: http.StatusOK,
		},
		"Succes POST": {
			method:     "POST",
			inputBody:  []byte(`{"login":"User","password":"Pass"}`),
			wantStatus: http.StatusOK,
		},
		"Wait BadRequest empty fields": {
			method:     "POST",
			inputBody:  []byte(`{"login":"","password":""}`),
			wantStatus: http.StatusBadRequest,
		},
		"Wait MethodNotAllowed": {
			method:     "ERROR",
			inputBody:  nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, srv.URL+"/signin", bytes.NewBuffer(test.inputBody))
			if err != nil {
				t.Errorf("Sign In request error: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Sign in response error: %v", err)
			}

			if resp.StatusCode != test.wantStatus {
				t.Errorf("Wait status: %v, but got: %v", test.wantStatus, resp.StatusCode)
			}
		})
	}
}
