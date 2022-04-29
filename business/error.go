package business

import "errors"

var (
	ErrUsernameRequired         = errors.New("username required")
	ErrUsernameAlreadyExisted   = errors.New("username already existed")
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrInvalidId                = errors.New("invalid id, id not numeric")
	ErrUserIdNotFound           = errors.New("user id not found")
	ErrEmailHasBeenRegister     = errors.New("email has been used")
	ErrUserId                   = errors.New("user id not found")
	ErrPasswordRequired         = errors.New("password is required")
	ErrEmailNotValid            = errors.New("email is not valid")
	ErrEmailRequired            = errors.New("email is required")
	ErrInvalidDate              = errors.New("invalid date, date must be formed : yyyy-mm-dd")
	ErrUsernamePasswordNotFound = errors.New("username or password empty")
	ErrInvalidAuthentication    = errors.New("authentication failed: invalid user credentials")
	ErrBadRequest               = errors.New("bad requests")
	ErrInvalidPayload           = errors.New("invalid payload")
	ErrNothingDestroy           = errors.New("no data found to delete")
	ErrInsufficientPermission   = errors.New("insufficient permission")
)
