package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		role := "user"
		if strings.HasPrefix(token, "admin_") {
			role = "admin"
		}

		c.Set("role", role)
		c.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
