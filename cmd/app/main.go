package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)

	router.Run("localhost:8080")
}
