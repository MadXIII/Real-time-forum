package errorface

import "errors"

type toSend struct {
	Err map[string]error
}

func New() *toSend {
	m := new(toSend)
	m.Err = make(map[string]error{
		"r": errors.New("asd"),
	})
}

func (t *toSend) Error(name string) string {
	if name == "ErrNotFound" {
		return "404 Not Found"
	}
	return ""
}

var (
	ErrNotFound       = errors.New("404 Not Found")
	ErrEmptyNickname  = errors.New("Nickname is empty")
	ErrEmptyEmail     = errors.New("Email is empty")
	ErrEmptyPassword  = errors.New("Password is empty")
	ErrEmptyConfirm   = errors.New("Confirm is empty")
	ErrEmptyFirstname = errors.New("Firstname is empty")
	ErrEmptyLastname  = errors.New("Lastname is empty")
	ErrEmptyGender    = errors.New("Gender is empty")
	ErrEmptyAge       = errors.New("Age is empty")
	ErrInvalidEmail   = errors.New("Invalid Email")
	ErrInvalidPass    = errors.New("Invalid Pass")
	ErrDiffSecondPass = errors.New("Different second password")
	ErrNoCookie       = errors.New("Can't find Cookie in store, to creat Post")
	ErrPostTitle      = errors.New("Title can't be empty and be more than 32 chars")
	ErrPostContent    = errors.New("Content can't be empty")
	ErrNilBody        = errors.New("Request Body doesn't be nil")
	ErrPassCompare    = errors.New("Password is uncomparable")
	ErrWrongPass      = errors.New("Wrong Password")
	ErrWrongLogin     = errors.New("Wrong Nickname or Email")
	ErrNickname       = errors.New("Nickname is already in use")
	ErrEmail          = errors.New("Email is already in use")
)
