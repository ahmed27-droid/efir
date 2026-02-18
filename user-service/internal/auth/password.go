package auth

import (
	"unicode"
	"user/internal/errs"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}


func ValidatePassword(password string) error {

	if len(password) < 8 {
		return errs.ErrPasswordTooShort
	}

	var hasLetter, hasDigit, hasSpecial bool

	for _, r := range password {
		switch {
		case unicode.IsLetter(r):
			hasLetter = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if !hasLetter || !hasDigit || !hasSpecial {
		return errs.ErrPasswordComplexity
	}
	return nil
}