package router

import (
	"Filmoteka/internal/controller"
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

func InitRoutes(cnt *controller.ActorController, mCnt *controller.MovieController) *gin.Engine {
	r := gin.Default()

	r.POST("/actor", cnt.CreateActor)
	r.PUT("/actor/update/:id", cnt.UpdateActor)
	r.GET("actor/get/:id", cnt.GetActor)
	r.DELETE("/actor/delete/:id", cnt.DeleteActor)

	r.POST("/movie", mCnt.CreateMovie)
	r.PUT("/movie/update/:id", mCnt.UpdateMovie)
	r.DELETE("/movie/delete/:id", mCnt.DeleteMovie)
	r.GET("/movie/get/:id", mCnt.GetMovie)

	return r
}

func (s Server) Run() error {
	return s.httpServer.ListenAndServe()

}
