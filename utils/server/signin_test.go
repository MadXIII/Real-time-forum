package server

import (
	"bytes"
	"fmt"
	"forum/utils/database/testdb"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
	"forum/utils/sessions/testsession"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	db := &testdb.TestDB{}
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
		// wantResult string
		wantError error
	}{

		// "Wait StatusOK  POST": {
		// 	method:     "POST",
		// 	password:   "password",
		// 	login:      "login",
		// 	wantStatus: http.StatusOK,
		// },
		// "GetUserByLogin error case": {
		// 	method:     "POST",
		// 	login:      "hunya",
		// 	password:   "1232132dl;askd",
		// 	wantStatus: http.StatusInternalServerError,
		// 	wantError:  newErr.ErrWrongLogin,
		// },
		// "GetUserByLogin good case": {
		// 	method:     "POST",
		// 	login:      "user",
		// 	password:   "123Password",
		// 	wantStatus: http.StatusOK,
		// },
		"checkLoginDatas error case": {
			method:     "POST",
			inputBody:  []byte(`{"login":"","password":""}`),
			wantStatus: http.StatusBadRequest,
			wantError:  newErr.ErrLoginData,
		},
		// "Wait MethodNotAllowed": {
		// 	method:     "ERROR",
		// 	inputBody:  nil,
		// 	wantStatus: http.StatusMethodNotAllowed,
		// },
		//{login:"user",password:"123Password"}
		// "Unmarshall error": {
		// 	method:     "POST",
		// 	wantStatus: http.StatusBadRequest,
		// 	wantError:  errors.New("invalid character 'l' looking for beginning of value"),
		// 	inputBody:  []byte(`ogin:"user",password:"123Password"}`),
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := bcrypt.GenerateFromPassword([]byte(test.password), bcrypt.MinCost)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println("Test", string(b))
			db.On("GetUserByLogin", test.login).Return(
				models.User{
					Nickname: test.login,
					Password: string(b),
				},
				test.wantError,
			)
			req, err := http.NewRequest(test.method, srv.URL+"/signin", bytes.NewBuffer(test.inputBody))
			fmt.Println("NewRequest", req)
			fmt.Println("inputBody", string(test.inputBody))
			//req, err := http.NewRequest(test.method, srv.URL+"/signin", bytes.NewBuffer(generateBody(test.password, test.login)))
			assert.Nil(t, err)
			resp, err := http.DefaultClient.Do(req)
			fmt.Println("Do", err)
			log.Print(resp, "inside test", err, test.wantError)

			if err != nil {
				assert.Equal(t, test.wantError, err.Error())
				assert.Equal(t, test.wantStatus, resp.StatusCode)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.wantStatus, resp.StatusCode)
			}

		})
	}
}

func generateBody(password, login string) []byte {
	return []byte(fmt.Sprintf(`{"login":"%s","password":"%s"}`, login, password))
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
