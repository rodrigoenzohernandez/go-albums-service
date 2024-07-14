package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/config"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/routes"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/utils/logger"
)

var log = logger.GetLogger("server")

func InitServer() {
	config.Load()

	apiHost := config.GetEnv("API_HOST", "0.0.0.0")
	apiPort := config.GetEnv("API_PORT", "8080")
	apiAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)

	db := repository.Connect()
	defer repository.Disconnect(db)
	albumRepo := repository.NewAlbumRepo(db)
	router := gin.Default()
	routes.RegisterAlbumRoutes(router, albumRepo)
	log.Info(fmt.Sprintf("Starting server on %s:%s", apiHost, apiPort))
	router.Run(apiAddress)
}
