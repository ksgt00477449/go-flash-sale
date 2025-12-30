package initialization

import (
	"go-flash-sale/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() (*gin.Engine, *gin.RouterGroup) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	corsMW := middleware.Cors()
	r.Use(corsMW)
	rg := r.Group("/api/v1")
	{
		rg.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
	return r, rg
}
