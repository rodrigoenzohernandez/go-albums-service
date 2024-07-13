package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/config"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/routes"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/database"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/utils/logger"
)

var log = logger.GetLogger("server")

func InitServer() {
	config.Load()
	db := database.Connect()
	defer database.Disconnect(db)
	router := gin.Default()
	routes.RegisterAlbumRoutes(router)
	log.Info("Starting server on localhost:8080")
	router.Run("localhost:8080")
}
