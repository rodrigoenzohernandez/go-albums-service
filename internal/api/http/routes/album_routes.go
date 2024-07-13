package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/handlers"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
)

func RegisterAlbumRoutes(router *gin.Engine, repo repository.AlbumRepositoryInterface) {
	albumHandler := handlers.NewAlbumHandler(repo)
	router.GET("/albums", albumHandler.GetAll)
}
