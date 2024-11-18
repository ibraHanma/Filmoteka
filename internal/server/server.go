package server

import (
	"Filmoteka/internal/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cnt    controller.ActorController
	mCnt   controller.MovieController
}

func NewServer(cnt controller.ActorController, mCnt controller.MovieController) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/actor", cnt.GetActor)
	engine.POST("/actor", cnt.CreateActor)
	engine.PUT("/actor/:id", cnt.UpdateActor)
	engine.DELETE("/actor/:id", cnt.DeleteActor)

	engine.GET("/movie", mCnt.GetMovie)
	engine.POST("/movie", mCnt.CreateMovie)
	engine.PUT("/movie/:id", mCnt.UpdateMovie)
	engine.DELETE("/movie/:id", mCnt.DeleteMovie)

	return &Server{
		engine: engine,
		cnt:    cnt,
		mCnt:   mCnt,
	}
}

func (s *Server) Run() error {

	return s.engine.Run(":8081")
}
