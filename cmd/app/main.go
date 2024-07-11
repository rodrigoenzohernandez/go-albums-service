package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterAlbumRoutes(router)

	router.Run("localhost:8080")
}
