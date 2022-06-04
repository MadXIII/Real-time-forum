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
	case ErrUnsignComment:
		return true
	case ErrUnsignVote:
		return true
	case ErrWrongCategory:
		return true
	default:
		return false
	}
}

var (
	ErrNotFound       = errors.New("404 Not Found")
	ErrEmptyNickname  = errors.New("nickname is empty")
	ErrEmptyEmail     = errors.New("email is empty")
	ErrEmptyPassword  = errors.New("password is empty")
	ErrEmptyConfirm   = errors.New("confirm is empty")
	ErrInvalidAge     = errors.New("age must be digtis, between 6-100")
	ErrInvalidEmail   = errors.New("invalid Email")
	ErrInvalidPass    = errors.New("invalid Pass")
	ErrDiffSecondPass = errors.New("different second password")
	ErrNoCookie       = errors.New("can't find Cookie in store, to creat Post")
	ErrPostTitle      = errors.New("title can't be empty and be more than 32 chars")
	ErrPostContent    = errors.New("content can't be empty")
	ErrNilBody        = errors.New("request Body doesn't be nil")
	ErrPassData       = errors.New("password can't be less than 8 chars and more than 32 chars")
	ErrLoginData      = errors.New("login field can't be empty")
	ErrWrongPass      = errors.New("wrong Password")
	ErrWrongLogin     = errors.New("wrong Nickname or Email")
	ErrNickname       = errors.New("nickname is already in use")
	ErrEmail          = errors.New("email is already in use")
	ErrDelCookie      = errors.New("can't delete cookie")
	ErrEmptyComment   = errors.New("comment is empty ")
	ErrLenComment     = errors.New("comment is too Long")
	ErrUnsignComment  = errors.New("need to login before to comment")
	ErrUnsignVote     = errors.New("need to login before to vote")
	ErrWrongCategory  = errors.New("wrong Category")
)
