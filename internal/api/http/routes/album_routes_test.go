package routes_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/api/http/routes"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AlbumRepositoryInterface struct {
	mock.Mock
}

func NewAlbumRepo(db *sql.DB) *repository.AlbumRepository {
	return &repository.AlbumRepository{DB: db}
}

func TestRegisterAlbumRoutes(t *testing.T) {
	db, _, _ := sqlmock.New()

	router := gin.Default()

	repo := NewAlbumRepo(db)

	routes.RegisterAlbumRoutes(router, repo)

	assert.Len(t, router.Routes(), 5, "Expected 5 routes to be registered")

	expectedRoutes := []struct {
		Method string
		Path   string
	}{
		{"GET", "/albums"},
		{"GET", "/albums/:id"},
		{"POST", "/albums"},
		{"PUT", "/albums/:id"},
		{"DELETE", "/albums/:id"},
	}

	for _, er := range expectedRoutes {
		found := false
		for _, r := range router.Routes() {
			if r.Method == er.Method && r.Path == er.Path {
				found = true
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("Route %s %s not found", er.Method, er.Path))
	}
}
