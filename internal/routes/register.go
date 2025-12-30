package routes

import (
	"go-flash-sale/internal/container"
	"go-flash-sale/internal/handler"
	"go-flash-sale/internal/initialization"
	"go-flash-sale/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(deps *container.Dependencies, middlewares *middleware.Middlewares) *gin.Engine {
	r := initialization.InitRoute()
	r.Use(middlewares.Cors)
	//  注册公共路由
	rg := r.Group("/api/v1")
	{
		rg.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
	// 业务路由注册
	userHandler := handler.NewUserHandler(deps.UserService)
	InitUserRoute(rg, userHandler, middlewares)
	return r
}
