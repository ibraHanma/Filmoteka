package server

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/middleware"
	"Filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.ActorController
	mCnt       controller.MovieController
}

func NewServer(cnt controller.ActorController, mCnt controller.MovieController) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			MaxHeaderBytes: 1 << 20,          // 1MB
			ReadTimeout:    10 * time.Second, // 10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt:  cnt,
		mCnt: mCnt,
	}
}
func (s Server) InitRoutes() {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	actorService := service.ActorService{}
	movieService := service.MovieService{}

	cnt := controller.ActorController{
		ActorService: actorService,
	}
	mCnt := controller.MovieController{MovieService: movieService}

	r.POST("/actor", cnt.CreateActor)
	r.PUT("/actor/update/:id", cnt.UpdateActor)
	r.DELETE("/actor/delete/:id", cnt.DeleteActor)
	r.GET("/actor/list", cnt.GetActor)

	r.POST("/movie", mCnt.CreateMovie)
	r.PUT("/movie/update/:id", mCnt.UpdateMovie)
	r.DELETE("/movie/delete/:id", mCnt.DeleteMovie)
	r.GET("/movie/list", mCnt.GetMovie)
	adminRoutes := r.Group("/admin")
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

func (s Server) Run() error {
	return s.httpServer.ListenAndServe()
}
func InitRoutes() *gin.Engine {
	router := gin.Default()

	actorController := &controller.ActorController{}

	// Определяем маршруты
	router.POST("/actors", actorController.CreateActor)

	return router
}
