package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		role, err := GetRoleFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("role", role)
		c.Next()
	}
}
func GetRoleFromToken(token string) (string, error) {

	if token == "valid_admin_token" {
		return "admin", nil
	} else if token == "valid_user_token" {
		return "user", nil
	}
	return "", fmt.Errorf("invalid token")
}
