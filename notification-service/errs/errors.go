package errs

import "errors"

var (
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrSubscriptionExists = errors.New("subscription already exists")
)
