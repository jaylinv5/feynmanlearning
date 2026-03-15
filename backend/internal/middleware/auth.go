package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jaylinv5/feynmanlearning/internal/model"
	"github.com/jaylinv5/feynmanlearning/internal/pkg/jwt"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Fail(401, "未提供认证令牌"))
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, model.Fail(401, "认证令牌格式错误"))
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Fail(401, "无效的认证令牌"))
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(int) != 3 {
			c.JSON(http.StatusForbidden, model.Fail(403, "需要管理员权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// TeacherMiddleware 教师权限中间件
func TeacherMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role.(int) != 2 && role.(int) != 3) {
			c.JSON(http.StatusForbidden, model.Fail(403, "需要教师或管理员权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetUserID 从上下文中获取用户ID
func GetUserID(c *gin.Context) uint64 {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint64)
}

// GetUserRole 从上下文中获取用户角色
func GetUserRole(c *gin.Context) int {
	role, exists := c.Get("role")
	if !exists {
		return 0
	}
	return role.(int)
}
