package server

import (
	"bytes"
	"errors"
	"forum/internal/database/testdb"
	newErr "forum/internal/error"
	"forum/internal/models"
	"forum/internal/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	db := &testdb.TestDB{}
	session := &testsession.TestSession{}
	mysrv := Init(db, session)
	mysrv.router.HandleFunc("/signin", mysrv.SignIn)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method     string
		signer     models.Sign
		inputBody  []byte
		wantStatus int
		wantError  error
	}{
		"Wait StatusOK": {
			method:     "POST",
			signer:     models.Sign{Login: "Name", Password: "123456Aa"},
			inputBody:  []byte(`{"login":"Name","password":"123456Aa"}`),
			wantStatus: http.StatusOK,
		},
		"GetUserByLogin error case": {
			method:     "POST",
			signer:     models.Sign{Login: "login", Password: "password"},
			inputBody:  []byte(`{"login":"login","password":"password"}`),
			wantStatus: http.StatusBadRequest,
			wantError:  newErr.ErrWrongLogin,
		},
		"checkLoginDatas error case": {
			method:     "POST",
			inputBody:  []byte(`{"login":"","password":""}`),
			wantStatus: http.StatusBadRequest,
			wantError:  newErr.ErrLoginData,
		},
		"Wait StatusMethodNotAllowed": {
			method:     "ERROR",
			inputBody:  nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
		"Unmarshall error case": {
			method:     "POST",
			wantStatus: http.StatusBadRequest,
			wantError:  errors.New("invalid character 'l' looking for beginning of value"),
			inputBody:  []byte(`login:"user",password:"123Password"}`),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := bcrypt.GenerateFromPassword([]byte(test.signer.Password), bcrypt.MinCost)
			if err != nil {
				t.Fatal(err)
			}

			db.On("GetUserByLogin", test.signer.Login).Return(
				models.User{
					Nickname: test.signer.Login,
					Password: string(b),
				},
				test.wantError,
			)
			session.On("CreateSession", 0).Return(&http.Cookie{
				Name:   "session",
				Value:  uuid.NewV4().String(),
				Path:   "/",
				MaxAge: 86400,
			})
			req, err := http.NewRequest(test.method, srv.URL+"/signin", bytes.NewBuffer(test.inputBody))
			assert.Nil(t, err)

			resp, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}

func TestCheckLoginDatas(t *testing.T) {
	tests := map[string]struct {
		inputData *models.Sign
		wantError error
	}{
		"Wait ErrLoginData if login field is empty": {
			inputData: &models.Sign{Login: ""},
			wantError: newErr.ErrLoginData,
		},
		"Wait ErrLoginData if login more than 32 chars": {
			inputData: &models.Sign{Login: string([]byte{32: '0'})},
			wantError: newErr.ErrLoginData,
		},
		"Wait ErrPassData if pass less than 8 chars": {
			inputData: &models.Sign{Login: "Login", Password: string([]byte{6, '0'})},
			wantError: newErr.ErrPassData,
		},
		"Wait ErrPassData if pass more than 32 chars": {
			inputData: &models.Sign{Login: "Login", Password: string([]byte{32: '0'})},
			wantError: newErr.ErrPassData,
		},
		"Success": {
			inputData: &models.Sign{Login: "Login", Password: "Password"},
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
