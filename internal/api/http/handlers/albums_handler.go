package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type Albums struct{}

// Returns all the albums as JSON.
func (a *Albums) GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Creates an album from JSON received in the request body.
func (a *Albums) Create(c *gin.Context) {
	var newAlbum album

	// Bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

// Returns a single album whose ID value matches the parameter.
func (a *Albums) GetByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
