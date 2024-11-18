package controller

import (
	"Filmoteka/internal/model"
	"Filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ServiceActor interface {
	CreateActor(name string, birthday time.Time, gender string) (int, error)
	GetActor(id int) (model.Actor, error)
	UpdateActor(id int, name string, birthday time.Time, gender string) error
	DeleteActor(id int) error
}

type ActorController struct {
	service      ServiceActor
	ActorService service.ActorService
}

func NewActorController(service ServiceActor) *ActorController {
	return &ActorController{service: service}
}
func (ac *ActorController) CreateActor(ctx *gin.Context) {
	var newActor model.Actor
	if err := ctx.ShouldBindJSON(&newActor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}
	if err := validate.StructExcept(newActor, "ID"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := ac.service.CreateActor(newActor.Name, newActor.Birthday, newActor.Gender)
	if err != nil {
		log.Printf("Error creating actor: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать актера"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (ac *ActorController) GetActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат идентификатора"})
		return
	}
	actor, err := ac.service.GetActor(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Актер не найден"})
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

func (ac *ActorController) UpdateActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат идентификатора"})
		return
	}

	var actor model.Actor
	if err := ctx.ShouldBindJSON(&actor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}
	actor.ID = id

	// Валидация структуры
	if err := validate.Struct(actor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.service.UpdateActor(id, actor.Name, actor.Birthday, actor.Gender); err != nil {
		log.Printf("Error updating actor: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось выполнить обновление"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Актер успешно обновлен"})
}

func (ac *ActorController) DeleteActor(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат идентификатора"})
		return
	}
	if err := ac.service.DeleteActor(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Актер не найден"})
		log.Printf("Error deleting actor with ID %d: %v", id, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Актер успешно удален"})
}
