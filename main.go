package main

import (
	"Filmoteka/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/public", PublicHandler)
	router.GET("/admin", middleware.TokenAuthMiddleware(), middleware.RequireRole("admin"), AdminHandler)
	router.Run(":8081")
}

func PublicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Public resource accessed!"})
}

func AdminHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin resource accessed! You have admin access."})
}
