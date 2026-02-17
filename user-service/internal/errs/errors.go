package errs

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrCheckEmailExists    = errors.New("failed to check email existence")
	ErrCheckUsernameExists = errors.New("failed to check username existence")
)
