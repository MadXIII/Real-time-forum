package server

import (
	"bytes"
	"fmt"
	"forum/database/testdb"
	newErr "forum/internal/error"
	"forum/models"
	"forum/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	db := &testdb.TestDB{}

	// db.InsertUser(models.User{Nickname: "testlogin", Password: string(byteses)})

	mysrv := Init(db, &testsession.TestSession{})
	mysrv.router.HandleFunc("/signin", mysrv.SignIn)
	srv := httptest.NewServer(&mysrv.router)

	// err := mysrv.DB

	tests := map[string]struct {
		method     string
		password   string
		login      string
		inputBody  []byte
		wantStatus int
	}{
		// "Wait StatusOK GET": {
		// 	method:     "GET",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusOK,
		// },
		"Wait StatusOK  POST": {
			method:     "POST",
			password:   "password",
			login:      "login",
			wantStatus: http.StatusOK,
		},
		// "Wait BadRequest empty fields": {
		// 	method:     "POST",
		// 	inputBody:  []byte(`{"login":"","password":""}`),
		// 	wantStatus: http.StatusBadRequest,
		// },
		// "Wait MethodNotAllowed": {
		// 	method:     "ERROR",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusMethodNotAllowed,
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := bcrypt.GenerateFromPassword([]byte(test.password), bcrypt.MinCost)
			if err != nil {
				t.Fatal(err)
			}

			db.On("GetUserByLogin", test.login).Return(
				&models.User{
					Nickname: test.login,
					Password: string(b),
				},
				nil,
			)

			req, err := http.NewRequest(test.method, srv.URL+"/signin", bytes.NewBuffer(generateBody(test.password, test.login)))
			assert.Nil(t, err)
			// if err != nil {
			// 	t.Errorf("Sign In request error: %v", err)
			// }
			resp, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)

			// if err != nil {
			// 	t.Errorf("Sign in response error: %v", err)
			// }
			assert.Equal(t, test.wantStatus, resp.StatusCode)
			// if resp.StatusCode != test.wantStatus {
			// 	t.Errorf("Wait status: %v, but got: %v", test.wantStatus, resp.StatusCode)
			// }
		})
	}
}

func generateBody(password, login string) []byte {
	return []byte(fmt.Sprintf(`{"login":"%s","password":"%s"}`, login, password))
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
