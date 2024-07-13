package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/handlers"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/middlewares"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
)

func RegisterAlbumRoutes(router *gin.Engine, repo repository.AlbumRepositoryInterface) {
	albumHandler := handlers.NewAlbumHandler(repo)
	IsValidAlbum := middlewares.IsValidAlbum(repo)

	router.GET("/albums", albumHandler.GetAll)
	router.GET("/albums/:id", albumHandler.GetByID)
	router.POST("/albums", IsValidAlbum, albumHandler.Create)
	router.PUT("/albums/:id", IsValidAlbum, albumHandler.Update)

}
