package server

import (
	"bytes"
	"forum/utils/database/testdb"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
	"forum/utils/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	mysrv := Init(&testdb.TestDB{}, &testsession.TestSession{})
	mysrv.router.HandleFunc("/signup", mysrv.SignUp)
	srv := httptest.NewServer(&mysrv.router)

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
		"Wait MethodNotAllowed": {
			method:     "ERROR",
			inputBody:  nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
		"Wait InternalServerError nil body": {
			method:     "POST",
			inputBody:  nil,
			wantStatus: http.StatusInternalServerError,
		},
		"Wait BadRequest empty fields": {
			method:     "POST",
			inputBody:  []byte(`{}`),
			wantStatus: http.StatusBadRequest,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, srv.URL+"/signup", bytes.NewBuffer(test.inputBody))
			if err != nil {
				t.Errorf("Sign Up request err: %v", err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Sign Up request err: %v", err)
			}
			if resp.StatusCode != test.wantStatus {
				t.Errorf("Wait status: %v, but got: %v", test.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestIsCorrectDatasToSignUp(t *testing.T) {
	tests := map[string]struct {
		inputDatas models.User
		wantError  error
	}{
		"Wait error with empty fields": {
			inputDatas: models.User{Nickname: ""},
			wantError:  newErr.ErrEmptyNickname,
		},
		"Wait error if email is not valid": {
			inputDatas: models.User{1, "nickname", "a@a.a", "password", "confirm", "firstname", "lastname", "gender", "7"},
			wantError:  newErr.ErrInvalidEmail,
		},
		"Wait error if confirm not same as password": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "password", "confirm", "firstname", "lastname", "gender", "7"},
			wantError:  newErr.ErrDiffSecondPass,
		},
		"Wait error if password is not valid": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "password", "password", "firstname", "lastname", "gender", "7"},
			wantError:  newErr.ErrInvalidPass,
		},
		"Success": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "123456Aa", "123456Aa", "firstname", "lastname", "gender", "7"},
			wantError:  nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isCorrectDatasToSignUp(test.inputDatas); err != test.wantError {
				t.Errorf("Wait for %v, but got %v", test.wantError, err)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := map[string]struct {
		inputEmail string
		wantErr    bool
	}{
		"Wait false if email was incorrect": {
			inputEmail: "a@a.a",
			wantErr:    false,
		},
		"Success": {
			inputEmail: "test@test.tt",
			wantErr:    true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isValidEmail(test.inputEmail); err != nil {
				t.Errorf("Wait for %v, but got %v", test.wantErr, err)
			}
		})
	}
}

func TestIsValidPass(t *testing.T) {
	tests := map[string]struct {
		inputPass string
		wantErr   bool
	}{
		"Wait false if pass less than 8 chars": {
			inputPass: string([]byte{6: '0'}),
			wantErr:   false,
		},
		"Wait false if pass more than 32 chars": {
			inputPass: string([]byte{32: '0'}),
			wantErr:   false,
		},
		"Wait false pass only latin chars": {
			inputPass: "123456Aaф",
			wantErr:   false,
		},
		"Wait false if pass has no Lower char": {
			inputPass: "123456AA",
			wantErr:   false,
		},
		"Wait false if pass has no Upper char": {
			inputPass: "123456aa",
			wantErr:   false,
		},
		"Wait false if pass has no Digit char": {
			inputPass: "AAAAaaaa",
			wantErr:   false,
		},
		"Success": {
			inputPass: "123456Aa",
			wantErr:   true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isValidPass(test.inputPass); err != nil {
				t.Errorf("Wait for %v, but got %v", test.wantErr, err)
			}
		})
	}
}

func TestCheckEmpty(t *testing.T) {
	tests := map[string]struct {
		inputUser models.User
		wantErr   error
	}{
		"Wait error with empty nickname": {
			inputUser: models.User{1, "", "email", "password", "confirm", "firstname", "lastname", "gender", "7"},
			wantErr:   newErr.ErrEmptyNickname,
		},
		"Wait error with empty email": {
			inputUser: models.User{1, "nickname", "", "password", "confirm", "firstname", "lastname", "gender", "7"},
			wantErr:   newErr.ErrEmptyEmail,
		},
		"Wait error with empty password": {
			inputUser: models.User{1, "nickname", "email", "", "confirm", "firstname", "lastname", "gender", "7"},
			wantErr:   newErr.ErrEmptyPassword,
		},
		"Wait error with empty cofirm": {
			inputUser: models.User{1, "nickname", "email", "password", "", "firstname", "lastname", "gender", "7"},
			wantErr:   newErr.ErrEmptyConfirm,
		},
		"Success": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "firstname", "lastname", "gender", "7"},
			wantErr:   nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := checkEmpty(test.inputUser); err != test.wantErr {
				t.Errorf("Wait for %v, but got %v", test.wantErr, err)
			}
		})
	}
}
