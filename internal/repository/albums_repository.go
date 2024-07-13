package repository

import (
	"database/sql"

	"github.com/rodrigoenzohernandez/web-service-gin/internal/models"
)

type album models.Album

type AlbumRepositoryInterface interface {
	SelectAll() ([]album, error)
}

type AlbumRepository struct {
	DB *sql.DB
}

func NewAlbumRepo(db *sql.DB) AlbumRepositoryInterface {
	return &AlbumRepository{DB: db}
}

func (repo *AlbumRepository) SelectAll() ([]album, error) {
	var albums []album

	query := `SELECT id, title, artist, price FROM "dev-schema".albums`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}
