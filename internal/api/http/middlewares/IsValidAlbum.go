package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/models"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
)

type Album models.Album

func IsValidAlbum(repo repository.AlbumRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		var album models.Album

		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Request Body", "error": err.Error()})
			c.Abort()
			return
		}

		exists, err := repo.AlbumExists(album.Title, album.Artist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to check album existence", "error": err.Error()})
			c.Abort()
			return
		}
		if exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The album already exists"})
			c.Abort()
			return
		}

		c.Set("album", album)

		c.Next()
	}
}
