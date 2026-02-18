package errors

import "errors"

var (
	ErrPostNotFound     = errors.New("post not found")
	ErrBroadcastService = errors.New("broadcast service error")
	ErrBroadcastNotActive = errors.New("broadcast is not active")
)
