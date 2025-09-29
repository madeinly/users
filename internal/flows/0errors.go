package flows

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrServerFailure      = errors.New("server issue check logs")
	ErrSessionExpired     = errors.New("session has expired")
)
