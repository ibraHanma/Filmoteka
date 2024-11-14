// middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware проверяет наличие токена и устанавливает роль в контекст
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Замените этот код на свою логику проверки аутентификации
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Здесь просто хардкодим роль для примера
		role := "user" // Вместо этого вы можете извлечь роль из своей базы данных или другого источника
		if strings.HasPrefix(token, "admin_") {
			role = "admin"
		}

		c.Set("role", role) // Устанавливаем роль в контекст
		c.Next()
	}
}

// RequireRole проверяет, есть ли у пользователя требуемая роль
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
