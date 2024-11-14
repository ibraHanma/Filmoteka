package main

import (
	"Filmoteka/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Открытые маршруты
	router.GET("/public", PublicHandler)

	// Защищенный маршрут для администраторов
	router.GET("/admin", middleware.TokenAuthMiddleware(), middleware.RequireRole("admin"), AdminHandler)

	router.Run(":8080")
	
}
func PublicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Public resource accessed!"})
}

func AdminHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin resource accessed! You have admin access."})
}
