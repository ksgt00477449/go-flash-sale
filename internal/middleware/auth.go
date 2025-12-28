package middleware

import (
	"go-flash-sale/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := extractToken(c)
	// token是否为空
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	// 验证token是否真正有效
	jwtService := auth.NewJWTService()
	claims, err := jwtService.ValidateToken(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	//查询reids中token是否存在
	tokenStore := auth.NewTokenStore()
	ok, err := tokenStore.Exists(c.Request.Context(), claims.ID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{"error": "Token has expired"})
		return
	}
	// 如果验证成功，可以将用户信息存储在上下文中，供后续处理使用
	c.Set("claims", claims) // claims 中包含用户id 用户名，用户email
	c.Next()
}

// extractToken 从请求中提取 JWT 字符串
func extractToken(c *gin.Context) string {
	// 1. 优先从 Authorization Header 中提取（Bearer Token）
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	// 2. 尝试从 Cookie 中提取（可选）
	if token, err := c.Cookie("token"); err == nil && token != "" {
		return token
	}

	// 3. 自定义请求头参数中提取（可选）
	if token := c.GetHeader("X-Auth-Token"); token != "" {
		return token
	}
	// 4. 都没找到，返回空字符串
	return ""
}
