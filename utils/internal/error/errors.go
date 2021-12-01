package error

import "errors"

func CheckErr(err error) bool {
	switch err {
	case ErrNotFound:
		return true
	case ErrEmptyNickname:
		return true
	case ErrEmptyEmail:
		return true
	case ErrEmptyPassword:
		return true
	case ErrEmptyConfirm:
		return true
	case ErrInvalidAge:
		return true
	case ErrInvalidEmail:
		return true
	case ErrInvalidPass:
		return true
	case ErrDiffSecondPass:
		return true
	case ErrNoCookie:
		return true
	case ErrPostTitle:
		return true
	case ErrPostContent:
		return true
	case ErrNilBody:
		return true
	case ErrPassData:
		return true
	case ErrLoginData:
		return true
	case ErrWrongPass:
		return true
	case ErrWrongLogin:
		return true
	case ErrNickname:
		return true
	case ErrEmail:
		return true
	case ErrDelCookie:
		return true
	case ErrEmptyComment:
		return true
	case ErrLenComment:
		return true
	case ErrUnsignVote:
		return true
	default:
		return false
	}
}

var (
	ErrNotFound       = errors.New("404 Not Found")
	ErrEmptyNickname  = errors.New("Nickname is empty")
	ErrEmptyEmail     = errors.New("Email is empty")
	ErrEmptyPassword  = errors.New("Password is empty")
	ErrEmptyConfirm   = errors.New("Confirm is empty")
	ErrInvalidAge     = errors.New("Age must be digtis, between 6-100")
	ErrInvalidEmail   = errors.New("Invalid Email")
	ErrInvalidPass    = errors.New("Invalid Pass")
	ErrDiffSecondPass = errors.New("Different second password")
	ErrNoCookie       = errors.New("Can't find Cookie in store, to creat Post")
	ErrPostTitle      = errors.New("Title can't be empty and be more than 32 chars")
	ErrPostContent    = errors.New("Content can't be empty")
	ErrNilBody        = errors.New("Request Body doesn't be nil")
	ErrPassData       = errors.New("Password can't be less than 8 chars and more than 32 chars")
	ErrLoginData      = errors.New("Login field can't be empty")
	ErrWrongPass      = errors.New("Wrong Password")
	ErrWrongLogin     = errors.New("Wrong Nickname or Email")
	ErrNickname       = errors.New("Nickname is already in use")
	ErrEmail          = errors.New("Email is already in use")
	ErrDelCookie      = errors.New("Can't delete cookie")
	ErrEmptyComment   = errors.New("Comment is empty ")
	ErrLenComment     = errors.New("Comment is too Long")
	ErrUnsignVote     = errors.New("Need to login befor vote")
)
