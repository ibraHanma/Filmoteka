package controller

import (
	"Filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Определяем интерфейс для MovieService
type serviceMovie interface {
	CreateMovie(title string, description string, releaseDate time.Time, rating int) (int, error)
	GetMovie(id int) (service.Movie, error)
	UpdateMovie(id int, title string, description string, releaseDate time.Time, rating int) (int, error)
	DeleteMovie(id int) (int, error)
}

// Структура Movie для передачи данных

type Movie struct {
	ID          int       `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	ReleaseDate time.Time `json:"ReleaseDate"`
	Rating      int       `json:"Rating"`
}

type MovieController struct {
	service      serviceMovie
	MovieService service.MovieService
}

func NewMovieController(service serviceMovie) *MovieController {
	return &MovieController{service: service}
}

// Создание нового фильма

func (mc *MovieController) CreateMovie(ctx *gin.Context) {
	var newMovie Movie
	if err := ctx.ShouldBindJSON(&newMovie); err != nil { // Используем ShouldBindJSON
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := mc.service.CreateMovie(newMovie.Title, newMovie.Description, newMovie.ReleaseDate, newMovie.Rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create movie"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"movie": id})
}

// Получение фильма по ID

func (mc *MovieController) GetMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	movie, err := mc.service.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Movie not found"})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

// Обновление фильма

func (mc *MovieController) UpdateMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var movie Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err = mc.service.UpdateMovie(id, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

// Удаление фильма

func (mc *MovieController) DeleteMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = mc.service.DeleteMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
