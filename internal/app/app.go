package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type App struct {
	r *gin.Engine
}

func NewApp() *App {
	deps, err := BuildDependencies()
	if err != nil {
		log.Fatalf("Failed to build dependencies: %v", err)
	}
	r := registerRoutes(deps)
	return &App{r: r}
}

func (a *App) Run() {
	go func() {
		log.Println("🚀 FlashSale server starting on :8080")
		if err := a.r.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
