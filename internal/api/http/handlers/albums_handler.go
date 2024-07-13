package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/models"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/utils/logger"
)

type Album models.Album

var log = logger.GetLogger("album_handler")

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

func (h *AlbumHandler) Create(c *gin.Context) {
	var album Album

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAlbum, err := h.Repo.Create(repository.Album(album))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)
}
