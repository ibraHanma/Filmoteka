package controller

import (
	"Filmoteka/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type serviceActor interface {
	CreateActor(name string, birthday time.Time, gender string) (int, error)
	GetActor(id int) (model.Actor, error)
	UpdateActor(id int, name string, birthday time.Time, gender string) error
	DeleteActor(id int) error
}

type ActorController struct {
	service serviceActor
}

func NewActorController(service serviceActor) *ActorController {
	return &ActorController{service: service}
}

func (ac *ActorController) CreateActor(ctx *gin.Context) {
	var newActor model.Actor
	if err := ctx.ShouldBindJSON(&newActor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, err := ac.service.CreateActor(newActor.Name, newActor.Birthday, newActor.Gender)
	{
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create actor"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"id": id})
	}

}

func (ac *ActorController) GetActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	actor, err := ac.service.GetActor(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor Not Found"})
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

func (ac *ActorController) UpdateActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var Actor model.Actor
	if err := ctx.ShouldBindJSON(&Actor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	Actor.ID = id

	if err := ac.service.UpdateActor(id, Actor.Name, Actor.Birthday, Actor.Gender); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor not Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Actor updated successfully"})
}

func (ac *ActorController) DeleteActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID Format"})
		return
	}
	if err := ac.service.DeleteActor(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor Not Found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)

}
