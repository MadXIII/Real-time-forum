package server

import (
	"bytes"
	"forum/database/testdb"
	newErr "forum/internal/error"
	"forum/models"
	"forum/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	db := &testdb.TestDB{}
	byteses, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	db.InsertUser(models.User{Nickname: "testlogin", Password: string(byteses)})

	mysrv := Init(db, &testsession.TestSession{})
	mysrv.router.HandleFunc("/signin", mysrv.SignIn)
	srv := httptest.NewServer(&mysrv.router)

	// err := mysrv.DB

	tests := map[string]struct {
		method     string
		inputBody  []byte
		wantStatus int
	}{
		"Wait StatusOK GET": {
			method:     "GET",
			inputBody:  nil,
			wantStatus: http.StatusOK,
		},
		"Wait StatusOK  POST": {
			method:     "POST",
			inputBody:  []byte(`{"login":"testlogin","password":"password"}`),
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

func TestCheckLoginDatas(t *testing.T) {
	tests := map[string]struct {
		inputData *Sign
		wantError error
	}{
		"Wait ErrLoginData if login field is empty": {
			inputData: &Sign{Login: ""},
			wantError: newErr.ErrLoginData,
		},
		"Wait ErrLoginData if login more than 32 chars": {
			inputData: &Sign{Login: string([]byte{32: '0'})},
			wantError: newErr.ErrLoginData,
		},
		"Wait ErrPassData if pass less than 8 chars": {
			inputData: &Sign{Login: "Login", Password: string([]byte{6, '0'})},
			wantError: newErr.ErrPassData,
		},
		"Wait ErrPassData if pass more than 32 chars": {
			inputData: &Sign{Login: "Login", Password: string([]byte{32: '0'})},
			wantError: newErr.ErrPassData,
		},
		"Success": {
			inputData: &Sign{Login: "Login", Password: "Password"},
			wantError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if result := checkLoginDatas(test.inputData); result != test.wantError {
				t.Errorf("Wait: %v, but got: %v", test.wantError, result)
			}
		})
	}
}
