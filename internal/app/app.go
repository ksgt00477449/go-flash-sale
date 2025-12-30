package app

import (
	"context"
	"go-flash-sale/internal/container"
	"go-flash-sale/internal/middleware"
	"go-flash-sale/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
	deps   *container.Dependencies
}

func NewApp() *App {
	deps, err := container.BuildDependencies()  //注册依赖
	mws := middleware.RegisterMiddlewares(deps) //注册中间件
	if err != nil {
		log.Fatalf("Failed to build dependencies: %v", err)
	}
	r := routes.RegisterRoutes(deps, mws) //注册路由
	return &App{
		server: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
		deps: deps}
}

func (a *App) Run() {

	// 启动 HTTP 服务（非阻塞）
	go func() {
		log.Println("🚀 Server starting on :8080")
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 监听中断信号（Ctrl+C 或 kill）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞等待信号

	log.Println("⏳ Shutting down server...")

	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务
	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	// 1. 关闭Redis
	err := a.deps.RedisClient.Close()
	if err != nil {
		log.Printf("Redis close error: %v", err)
	}
	// 2. 关闭数据库
	if sqlDB, err := a.deps.DB.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Database close error: %v", err)
		}
	}
	log.Println("✅ Server exited gracefully")

}
