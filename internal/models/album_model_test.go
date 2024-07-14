package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAlbumModel(t *testing.T) {
	tests := []struct {
		name   string
		album  Album
		expect string
	}{
		{"Missing Title", Album{ID: uuid.NewString(), Artist: "Acru", Price: 9.99}, "Title"},
		{"Missing Artist", Album{ID: uuid.NewString(), Title: "La balada del diablo y la muerte", Price: 19.99}, "Artist"},
		{"Missing Price", Album{ID: uuid.NewString(), Title: "Todo de oro", Artist: "YSY A"}, "Price"},
		{"All", Album{ID: uuid.NewString(), Title: "0800", Artist: "Veeyamn", Price: 5.99}, "All"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.expect {
			case "Title":
				assert.Empty(t, tc.album.Title, "Expected Title to be missing")
			case "Artist":
				assert.Empty(t, tc.album.Artist, "Expected Artist to be missing")
			case "Price":
				assert.Zero(t, tc.album.Price, "Expected Price to be 0")
			case "All":
				assert.NotEmpty(t, tc.album.ID, "ID should not be empty")
				assert.NotEmpty(t, tc.album.Title, "Title should not be empty")
				assert.NotEmpty(t, tc.album.Artist, "Artist should not be empty")
				assert.NotZero(t, tc.album.Price, "Price should not be 0")
			}

		})
	}
}
