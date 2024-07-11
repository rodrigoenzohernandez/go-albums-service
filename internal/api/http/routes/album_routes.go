package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/handlers"
)

var albumHandler = handlers.Albums{}

func RegisterAlbumRoutes(router *gin.Engine) {
	router.GET("/albums", albumHandler.GetAll)
	router.GET("/albums/:id", albumHandler.GetByID)
	router.POST("/albums", albumHandler.Create)
}
