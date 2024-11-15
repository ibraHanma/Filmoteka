package controller

import (
	"Filmoteka/internal/model"
	"Filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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
	service      serviceActor
	ActorService service.ActorService
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create actor"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
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

	var actor model.Actor
	if err := ctx.ShouldBindJSON(&actor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	actor.ID = id

	if err := ac.service.UpdateActor(id, actor.Name, actor.Birthday, actor.Gender); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Actor updated successfully"})
}
func (ac *ActorController) DeleteActor(ctx *gin.Context) {
	if ac.service == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Service not initialized"})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID Format"})
		return
	}
	if err := ac.service.DeleteActor(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor Not Found"})
		log.Printf("Error deleting actor with ID %d: %v", id, err) // Логирование ошибки
		return
	}

}
