package errors

import "errors"

// ==============token相关错误定义==================
var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrTokenNotActive    = errors.New("token not active yet")
	ErrCreateTokenFailed = errors.New("failed to create token")
)
