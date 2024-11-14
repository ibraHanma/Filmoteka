package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type ContextKey string

const UserRoleKey ContextKey = "userRole"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Ожидаем формат "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		role, err := parseToken(token) // Функция для парсинга токена
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Устанавливаем роль в контекст
		ctx := context.WithValue(r.Context(), UserRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// parseToken - функция для парсинга токена и получения роли пользователя
// Здесь вы можете добавить вашу логику для проверки токена
func parseToken(token string) (string, error) {
	// Пример: просто возвращаем "admin" для демонстрации
	// В реальном приложении вы должны проверить токен и вернуть соответствующую роль
	if token == "your-secret-token" {
		return "admin", nil
	}
	return "", fmt.Errorf("invalid token")
}
