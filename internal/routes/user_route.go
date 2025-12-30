package routes

import (
	"go-flash-sale/internal/handler"
	"go-flash-sale/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(r *gin.RouterGroup, userHandler *handler.UserHandler, middleware *middleware.Middlewares) {
	rg := r.Group("/user")
	rg.POST("/register", userHandler.Register)
	rg.POST("/login", userHandler.Login)
}
