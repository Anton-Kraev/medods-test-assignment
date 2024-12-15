package errs

import "errors"

var (
	ErrSessionNotFound = errors.New("session for user not found")

	ErrInvalidRefreshToken = errors.New("refresh token invalid")
)
