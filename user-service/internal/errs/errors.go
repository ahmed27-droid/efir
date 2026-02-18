package errs

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrCheckEmailExists      = errors.New("failed to check email existence")
	ErrCheckUsernameExists   = errors.New("failed to check username existence")
	ErrPasswordTooShort      = errors.New("password must be at least 8 characters")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrPasswordComplexity    = errors.New("password must contain letter, digit and special character")
)
