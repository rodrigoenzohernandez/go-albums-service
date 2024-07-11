package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/handlers"
)

func main() {
	router := gin.Default()
	albumHandler := handlers.Albums{}

	router.GET("/albums", albumHandler.GetAll)
	router.GET("/albums/:id", albumHandler.GetByID)
	router.POST("/albums", albumHandler.Create)

	router.Run("localhost:8080")
}
