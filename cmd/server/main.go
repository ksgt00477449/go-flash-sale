package main

import (
	"go-flash-sale/internal/cache"
	"go-flash-sale/internal/handler"
	"go-flash-sale/internal/initialization"
	"go-flash-sale/internal/middleware"
	"go-flash-sale/internal/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 初始化路由 并设置全局跨域问题中间件 全局路由前缀/api/v1
	r, rg := initialization.InitRoute()
	// 初始化数据库链接
	db := initialization.InitDB()
	// 初始化Redis链接
	redisClient := initialization.InitRedis()
	// 自动迁移模式，创建或更新表结构
	initialization.InitTableAutoMigrate(db)
	// 初始化依赖
	tokenCache := cache.NewTokenCache(redisClient)
	_ = tokenCache // 防止未使用警告，后续可删除
	// 初始化中间件
	authMW := middleware.AuthMiddleware(redisClient)
	_ = authMW // 防止未使用警告，后续可删除

	// 注册handler
	userHandler := handler.NewUserHandler(db, redisClient)

	// 业务路由注册
	routes.InitUserRoute(rg, userHandler)

	go func() {
		log.Println("🚀 FlashSale server starting on :8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
