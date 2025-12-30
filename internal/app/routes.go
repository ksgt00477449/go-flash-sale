package app

import (
	"go-flash-sale/internal/handler"
	"go-flash-sale/internal/initialization"
	"go-flash-sale/internal/middleware"
	"go-flash-sale/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(deps *Dependencies) *gin.Engine {
	r := initialization.InitRoute()
	userHandler := handler.NewUserHandler(deps.UserService)
	authMW := middleware.AuthMiddleware(deps.RedisClient)
	_ = authMW // 防止未使用警告，后续可删除
	// 全局中间件
	corsMW := middleware.Cors()
	r.Use(corsMW)
	//  注册公共路由
	rg := r.Group("/api/v1")
	{
		rg.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
	// 业务路由注册
	routes.InitUserRoute(rg, userHandler)
	return r
}
