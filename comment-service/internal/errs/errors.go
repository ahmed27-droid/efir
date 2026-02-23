package errs

import "errors"

var (
	ErrReactionNotFound     = errors.New("reaction not found")
	ErrCommentNotFound      = errors.New("comment not found")
	ErrPostNotFound         = errors.New("post not found")
	ErrBroadcastService     = errors.New("broadcast service error")
	ErrBroadcastNotActive   = errors.New("broadcast is not active")
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
)
