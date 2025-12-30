package initialization

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	return r
}
