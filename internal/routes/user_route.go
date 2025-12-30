package routes

import (
	"go-flash-sale/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	rg := r.Group("/user")
	rg.POST("/register", userHandler.Register)
	rg.POST("/login", userHandler.Login)
}
