package repository

import (
	"database/sql"

	"github.com/rodrigoenzohernandez/web-service-gin/internal/models"
)

type album models.Album

type AlbumRepositoryInterface interface {
	SelectAll() ([]album, error)
	SelectByID(id string) (*album, error)
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

func (repo *AlbumRepository) SelectByID(id string) (*album, error) {
	var a album
	query := `SELECT id, title, artist, price FROM "dev-schema".albums WHERE id = $1`

	err := repo.DB.QueryRow(query, id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return &a, nil
}
