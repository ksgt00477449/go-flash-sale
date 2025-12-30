package middleware

import (
	"go-flash-sale/internal/container"

	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	Auth gin.HandlerFunc
	Cors gin.HandlerFunc
	// RateLimit  gin.HandlerFunc
	// Logging    gin.HandlerFunc
	// Tracing    gin.HandlerFunc
	// Permission gin.HandlerFunc
	// 可按需扩展
}

func RegisterMiddlewares(deps *container.Dependencies) *Middlewares {
	auth := AuthMiddleware(deps.RedisClient)
	cors := CorsMiddleware()
	return &Middlewares{
		Auth: auth,
		Cors: cors,
		// RateLimit:  rateLimit,
		// Logging:    logging,
		// Tracing:    tracing,
		// Permission: permission,
	}
}
