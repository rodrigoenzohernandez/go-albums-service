package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rodrigoenzohernandez/go-albums-service/internal/models"
	"github.com/rodrigoenzohernandez/go-albums-service/internal/repository"
	"github.com/rodrigoenzohernandez/go-albums-service/internal/utils/logger"
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
	albumSetInMiddleware, _ := c.Get("album")
	album, _ := albumSetInMiddleware.(models.Album)

	createdAlbum, err := h.Repo.Create(repository.Album(album))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating an album": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)
}

func (h *AlbumHandler) Update(c *gin.Context) {
	album := getAlbumFromContext(c)

	id := c.Param("id")

	isUUID := isUUID(c, id)
	if !isUUID {
		return
	}

	updatedAlbum, err := h.Repo.Update(id, repository.Album(album))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error updating an album": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAlbum)
}

// Deletes an album by ID.
func (h *AlbumHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	isUUID := isUUID(c, id)
	if !isUUID {
		return
	}

	err := h.Repo.Delete(id)
	if err != nil {
		const notFoundMsg = "Album not found"
		if err.Error() == notFoundMsg {
			c.JSON(http.StatusNotFound, gin.H{"msg": notFoundMsg})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting an album"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

// Returns the album from the context. This album is set in the middleware.
func getAlbumFromContext(c *gin.Context) models.Album {
	albumSetInMiddleware, _ := c.Get("album")
	album, _ := albumSetInMiddleware.(models.Album)
	return album
}

func isUUID(c *gin.Context, text string) bool {

	if _, err := uuid.Parse(text); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format, UUID is expected"})
		return false
	}
	return true
}
