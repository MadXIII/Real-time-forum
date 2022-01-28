package server

import (
	"bytes"
	"forum/internal/database/testdb"
	newErr "forum/internal/error"
	"forum/internal/models"
	"forum/internal/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	db := &testdb.TestDB{}
	mysrv := Init(db, &testsession.TestSession{})
	mysrv.router.HandleFunc("/signup", mysrv.SignUp)
	srv := httptest.NewServer(&mysrv.router)

	tests := map[string]struct {
		method     string
		user       models.User
		inputBody  []byte
		wantStatus int
		wantError  error
	}{
		"Wait StatusOK": {
			method:     "POST",
			user:       models.User{ID: 1, Nickname: "Nick", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
			inputBody:  []byte(`{"id":1,"nickname":"Nick","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusOK,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := bcrypt.GenerateFromPassword([]byte(test.user.Password), bcrypt.MinCost)
			if err != nil {
				t.Fatal(err)
			}
			test.user.Password = string(b)

			db.On("InsertUser", &test.user).Return(test.wantError)

			req, err := http.NewRequest(test.method, srv.URL+"/signup", bytes.NewBuffer(test.inputBody))
			assert.Nil(t, err)
			resp, err := http.DefaultClient.Do(req)

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
		"Wait error if age is not valid": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "123456Aa", "123456Aa", "firstname", "lastname", "gender", "seven"},
			wantError:  newErr.ErrInvalidAge,
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
		wantErr    error
	}{
		"Wait false if email was incorrect": {
			inputEmail: "a@a.a",
			wantErr:    newErr.ErrInvalidEmail,
		},
		"Success": {
			inputEmail: "test@test.tt",
			wantErr:    nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isValidEmail(test.inputEmail); err != test.wantErr {
				t.Errorf("Wait for %v, but got %v", test.wantErr, err)
			}
		})
	}
}

func TestIsValidPass(t *testing.T) {
	tests := map[string]struct {
		inputPass string
		wantErr   error
	}{
		"Wait error if pass less than 8 chars": {
			inputPass: string([]byte{6: '0'}),
			wantErr:   newErr.ErrInvalidPass,
		},
		"Wait error if pass more than 32 chars": {
			inputPass: string([]byte{32: '0'}),
			wantErr:   newErr.ErrInvalidPass,
		},
		"Wait error pass only latin chars": {
			inputPass: "123456Aaф",
			wantErr:   newErr.ErrInvalidPass,
		},
		"Wait error if pass has no Lower char": {
			inputPass: "123456AA",
			wantErr:   newErr.ErrInvalidPass,
		},
		"Wait error if pass has no Upper char": {
			inputPass: "123456aa",
			wantErr:   newErr.ErrInvalidPass,
		},
		"Wait error if pass has no Digit char": {
			inputPass: "AAAAaaaa",
			wantErr:   newErr.ErrInvalidPass,
		},
		"Success": {
			inputPass: "123456Aa",
			wantErr:   nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isValidPass(test.inputPass); err != test.wantErr {
				t.Errorf("Wait for %v, but got %v", test.wantErr, err)
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	tests := map[string]struct {
		inputAge string
		wantErr  error
	}{
		"Wait error if age not number": {
			inputAge: "age",
			wantErr:  newErr.ErrInvalidAge,
		},
		"Wait error if user too young": {
			inputAge: "5",
			wantErr:  newErr.ErrInvalidAge,
		},
		"Wait error if user too old": {
			inputAge: "101",
			wantErr:  newErr.ErrInvalidAge,
		},
		"Success with age": {
			inputAge: "22",
			wantErr:  nil,
		},
		"Success with empty field": {
			inputAge: "",
			wantErr:  nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if err := isValidAge(test.inputAge); err != test.wantErr {
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
