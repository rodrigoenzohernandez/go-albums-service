package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/models"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/utils/logger"
)

type album models.Album

var log = logger.GetLogger("album_handler")

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type Albums struct{}

type AlbumHandler struct {
	Repo repository.AlbumRepositoryInterface
}

func NewAlbumHandler(repo repository.AlbumRepositoryInterface) *AlbumHandler {
	return &AlbumHandler{Repo: repo}
}

// Returns all the albums as JSON.
func (h *AlbumHandler) GetAll(c *gin.Context) {
	albums, err := h.Repo.SelectAll()
	if err != nil {
		log.Error("Error on GetAll")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// Returns an album by ID as JSON.
func (h *AlbumHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format, UUID is expected"})
		return
	}

	album, err := h.Repo.SelectByID(id)
	if err != nil {
		log.Error("Error on GetByID")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "error": err.Error()})
		return
	}

	if album == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// Creates an album from JSON received in the request body.
func (a *Albums) Create(c *gin.Context) {
	var newAlbum album

	// Bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Error("Error on Album creation")
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
