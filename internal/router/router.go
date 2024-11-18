package router

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	cnt    controller.ActorController
	mCnt   controller.MovieController
}

func (s *Server) InitRoutes() {

	s.engine.Use(middleware.AuthMiddleware())

	adminRoutes := s.engine.Group("/admin")
	adminRoutes.Use(AdminMiddleware())

	adminRoutes.POST("/actor", s.cnt.CreateActor)
	adminRoutes.POST("/movie", s.mCnt.CreateMovie)
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: insufficient permissions"})
			c.Abort()
			return
		}
		c.Next()
	}
}
