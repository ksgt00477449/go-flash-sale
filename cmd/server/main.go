package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置 Gin 为 release 模式（关闭调试日志）
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery()) // 自动恢复 panic

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "flashsale-go",
		})
	})

	log.Println("🚀 FlashSale server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
