package server

import (
	"bytes"
	"errors"
	"forum/internal/database"
	"forum/internal/database/testdb"
	newErr "forum/internal/error"
	"forum/internal/models"
	s "forum/internal/sessions"
	"forum/internal/sessions/testsession"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	type args struct {
		user      models.User
		wantError error
	}

	tests := []struct {
		name       string
		args       args
		method     string
		inputBody  []byte
		wantStatus int
		callback   func(args) (database.Repository, s.Repository)
	}{
		{
			name: "Wait StatusOk",
			args: args{
				user: models.User{ID: 1, Nickname: "Nick", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"nickname":"Nick","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusOK,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("InsertUser", mock.Anything).Return(a.wantError)
				session.On("CreateSession", a.user.ID).Return(&http.Cookie{Name: "session"})
				return db, session
			},
		},
		{
			name:       "Wait StatusMethodNotAllowed status",
			args:       args{},
			method:     http.MethodPatch,
			wantStatus: http.StatusMethodNotAllowed,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				return db, session
			},
		},
		{
			name: "Wait Unmarshal error",
			args: args{
				user: models.User{ID: 1, Nickname: "Nick", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
			},
			method:     http.MethodPost,
			inputBody:  []byte(`"id":1,"nickname":"Nick","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				return db, session
			},
		},
		{
			name: "Wait isCorrectDatasToSignUp error",
			args: args{
				user: models.User{ID: 1, Nickname: "", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"nickname":"","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("InsertUser", mock.Anything).Return(a.wantError)
				session.On("CreateSession", a.user.ID).Return(&http.Cookie{Name: "session"})
				return db, session
			},
		},
		{
			name: "Wait ErrNickname error",
			args: args{
				user:      models.User{ID: 1, Nickname: "Nick", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
				wantError: errors.New("InsertUser, Exec: UNIQUE constraint failed: user.nickname"),
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"nickname":"Nick","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("InsertUser", mock.Anything).Return(a.wantError)
				return db, session
			},
		},
		{
			name: "Wait ErrEmail error",
			args: args{
				user:      models.User{ID: 1, Nickname: "Nick", Email: "test@test.tt", Password: "123456Aa", Confirm: "123456Aa"},
				wantError: errors.New("InsertUser, Exec: UNIQUE constraint failed: user.email"),
			},
			method:     http.MethodPost,
			inputBody:  []byte(`{"id":1,"nickname":"Nick","email":"test@test.tt","password":"123456Aa","confirm":"123456Aa"}`),
			wantStatus: http.StatusBadRequest,
			callback: func(a args) (database.Repository, s.Repository) {
				db := &testdb.TestDB{}
				session := &testsession.TestSession{}
				db.On("InsertUser", mock.Anything).Return(a.wantError)
				return db, session
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, session := test.callback(test.args)
			mysrv := Init(db, session)

			mysrv.router.HandleFunc("/signup", mysrv.SignUp)

			req, err := http.NewRequest(test.method, "/signup", bytes.NewBuffer(test.inputBody))
			assert.Nil(t, err)

			recorder := httptest.NewRecorder()
			mysrv.SignUp(recorder, req)

			resp := recorder.Result()

			assert.Equal(t, test.wantStatus, resp.StatusCode)
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
