package auth

import "errors"

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrTokenNotActive    = errors.New("token not active yet")
	ErrCreateTokenFailed = errors.New("failed to create token")
)
