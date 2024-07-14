package repository

import (
	"database/sql"
	"errors"

	"github.com/rodrigoenzohernandez/go-albums-service/internal/models"
)

type Album models.Album

type AlbumRepositoryInterface interface {
	SelectAll() ([]Album, error)
	SelectByID(id string) (*Album, error)
	Create(album Album) (*Album, error)
	AlbumExists(title, artist string) (bool, error)
	Update(id string, album Album) (*Album, error)
	Delete(id string) error
}

type AlbumRepository struct {
	DB *sql.DB
}

func NewAlbumRepo(db *sql.DB) AlbumRepositoryInterface {
	return &AlbumRepository{DB: db}
}

func (repo *AlbumRepository) SelectAll() ([]Album, error) {
	var albums []Album

	query := `SELECT id, title, artist, price FROM "dev-schema".albums`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
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

func (repo *AlbumRepository) SelectByID(id string) (*Album, error) {
	var a Album
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

func (repo *AlbumRepository) Create(album Album) (*Album, error) {
	query := `INSERT INTO "dev-schema".albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id`

	err := repo.DB.QueryRow(query, album.Title, album.Artist, album.Price).Scan(&album.ID)
	if err != nil {
		return &Album{}, err
	}

	return &album, nil
}

func (repo *AlbumRepository) AlbumExists(title, artist string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM "dev-schema".albums WHERE title = $1 AND artist = $2)`
	var exists bool
	err := repo.DB.QueryRow(query, title, artist).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo *AlbumRepository) Update(id string, album Album) (*Album, error) {
	query := `UPDATE "dev-schema".albums SET title = $1, artist = $2, price = $3 WHERE id = $4 RETURNING id, title, artist, price`

	err := repo.DB.QueryRow(query, album.Title, album.Artist, album.Price, id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (repo *AlbumRepository) Delete(id string) error {
	query := `DELETE FROM "dev-schema".albums WHERE id = $1`

	result, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Album not found")
	}

	return nil
}
