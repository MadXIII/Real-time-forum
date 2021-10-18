package server

import (
	newErr "forum/internal/error"
	"forum/models"
	"testing"
)

func TestIsValidPass(t *testing.T) {
	tests := map[string]struct {
		inputPass  string
		wantResult bool
	}{
		"Wait false if pass less than 8 chars": {
			inputPass:  "1234567",
			wantResult: false,
		},
		"Wait false if pass more than 32 chars" : {
			inputPass:
		}
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
