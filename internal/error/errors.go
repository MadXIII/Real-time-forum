package error

import "errors"

var (
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
)
