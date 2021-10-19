package server

import (
	newErr "forum/internal/error"
	"forum/models"
	"testing"
)

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
			inputDatas: models.User{1, "nickname", "a@a.a", "password", "confirm", "firstname", "lastname", "gender", 7},
			wantError:  newErr.ErrInvalidEmail,
		},
		"Wait error if confirm not same as password": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "password", "confirm", "firstname", "lastname", "gender", 7},
			wantError:  newErr.ErrDiffSecondPass,
		},
		"Wait error if password is not valid": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "password", "password", "firstname", "lastname", "gender", 7},
			wantError:  newErr.ErrInvalidPass,
		},
		"Success": {
			inputDatas: models.User{1, "nickname", "mail@mail.ru", "123456Aa", "123456Aa", "firstname", "lastname", "gender", 7},
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
		wantResult bool
	}{
		"Wait false if email was incorrect": {
			inputEmail: "a@a.a",
			wantResult: false,
		},
		"Success": {
			inputEmail: "test@test.tt",
			wantResult: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if result := isValidEmail(test.inputEmail); result != test.wantResult {
				t.Errorf("Wait for %v, but got %v", test.wantResult, result)
			}
		})
	}
}

func TestIsValidPass(t *testing.T) {
	tests := map[string]struct {
		inputPass  string
		wantResult bool
	}{
		"Wait false if pass less than 8 chars": {
			inputPass:  string([]byte{6: '0'}),
			wantResult: false,
		},
		"Wait false if pass more than 32 chars": {
			inputPass:  string([]byte{32: '0'}),
			wantResult: false,
		},
		"Wait false if pass has no Lower char": {
			inputPass:  "123456AA",
			wantResult: false,
		},
		"Wait false if pass has no Upper char": {
			inputPass:  "123456aa",
			wantResult: false,
		},
		"Wait false if pass has no Digit char": {
			inputPass:  "AAAAaaaa",
			wantResult: false,
		},
		"Success": {
			inputPass:  "123456Aa",
			wantResult: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if result := isValidPass(test.inputPass); result != test.wantResult {
				t.Errorf("Wait for %v, but got %v", test.wantResult, result)
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
			inputUser: models.User{1, "", "email", "password", "confirm", "firstname", "lastname", "gender", 7},
			wantErr:   newErr.ErrEmptyNickname,
		},
		"Wait error with empty email": {
			inputUser: models.User{1, "nickname", "", "password", "confirm", "firstname", "lastname", "gender", 7},
			wantErr:   newErr.ErrEmptyEmail,
		},
		"Wait error with empty password": {
			inputUser: models.User{1, "nickname", "email", "", "confirm", "firstname", "lastname", "gender", 7},
			wantErr:   newErr.ErrEmptyPassword,
		},
		"Wait error with empty cofirm": {
			inputUser: models.User{1, "nickname", "email", "password", "", "firstname", "lastname", "gender", 7},
			wantErr:   newErr.ErrEmptyConfirm,
		},
		"Wait error with empty firstname": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "", "lastname", "gender", 7},
			wantErr:   newErr.ErrEmptyFirstname,
		},
		"Wait error with empty lastname": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "firstname", "", "gender", 7},
			wantErr:   newErr.ErrEmptyLastname,
		},
		"Wait error with empty gender": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "firstname", "lastname", "", 7},
			wantErr:   newErr.ErrEmptyGender,
		},
		"Wait error with empty age": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "firstname", "lastname", "gender", 0},
			wantErr:   newErr.ErrEmptyAge,
		},
		"Success": {
			inputUser: models.User{1, "nickname", "email", "password", "confirm", "firstname", "lastname", "gender", 7},
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
