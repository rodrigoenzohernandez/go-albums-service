package repository_tests

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodrigoenzohernandez/go-albums-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func NewAlbumRepo(db *sql.DB) *repository.AlbumRepository {
	return &repository.AlbumRepository{DB: db}
}

func InitMocks(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *repository.AlbumRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() { db.Close() })

	repo := NewAlbumRepo(db)
	return db, mock, repo
}

func TestSelectAll(t *testing.T) {
	_, mock, repo := InitMocks(t)

	rows := sqlmock.NewRows([]string{"id", "title", "artist", "price"}).
		AddRow("1", "Blue Train", "John Coltrane", 56.99).
		AddRow("2", "Giant Steps", "John Coltrane", 63.99)

	mock.ExpectQuery("^SELECT id, title, artist, price FROM \"dev-schema\"\\.albums$").WillReturnRows(rows)

	albums, err := repo.SelectAll()
	assert.NoError(t, err)
	assert.Len(t, albums, 2)
	assert.Equal(t, "Blue Train", albums[0].Title)
	assert.Equal(t, "John Coltrane", albums[1].Artist)
	assert.NotEqual(t, "Eminem", albums[1].Artist)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestSelectByID(t *testing.T) {
	_, mock, repo := InitMocks(t)

	rows := sqlmock.NewRows([]string{"id", "title", "artist", "price"}).
		AddRow("1", "Blue Train", "John Coltrane", 56.99)

	mock.ExpectQuery(`SELECT id, title, artist, price FROM "dev-schema"\.albums WHERE id = \$1`).WillReturnRows(rows)
	album, err := repo.SelectByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "Blue Train", album.Title)
	assert.Equal(t, "John Coltrane", album.Artist)
	assert.NotEqual(t, "Eminem", album.Artist)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestAlbumExists(t *testing.T) {
	_, mock, repo := InitMocks(t)

	rows := sqlmock.NewRows([]string{"exists"}).
		AddRow(true)

	mock.ExpectQuery(`SELECT EXISTS\(SELECT 1 FROM "dev-schema"\.albums WHERE title = \$1 AND artist = \$2\)`).
		WithArgs("Blue Train", "John Coltrane").
		WillReturnRows(rows)

	exists, err := repo.AlbumExists("Blue Train", "John Coltrane")
	assert.NoError(t, err)
	assert.True(t, exists)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDelete(t *testing.T) {
	_, mock, repo := InitMocks(t)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM \"dev-schema\".albums WHERE id = \\$1").
			WithArgs("existing-id").
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Delete("existing-id")
		assert.NoError(t, err)
	})

	t.Run("Album not found", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM \"dev-schema\".albums WHERE id = \\$1").
			WithArgs("non-existing-id").
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete("non-existing-id")
		assert.EqualError(t, err, "Album not found")
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM \"dev-schema\".albums WHERE id = \\$1").
			WithArgs("any-id").
			WillReturnError(errors.New("db error"))

		err := repo.Delete("any-id")
		assert.EqualError(t, err, "db error")
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	_, mock, repo := InitMocks(t)

	newAlbum := repository.Album{Title: "New Album", Artist: "New Artist", Price: 9.99}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery("INSERT INTO \"dev-schema\".albums").
			WithArgs(newAlbum.Title, newAlbum.Artist, newAlbum.Price).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("new-id"))

		createdAlbum, err := repo.Create(newAlbum)

		assert.NoError(t, err)
		assert.NotNil(t, createdAlbum)
		assert.Equal(t, "new-id", createdAlbum.ID)
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery("INSERT INTO \"dev-schema\".albums").
			WithArgs(newAlbum.Title, newAlbum.Artist, newAlbum.Price).
			WillReturnError(errors.New("db error"))

		_, err := repo.Create(newAlbum)

		assert.EqualError(t, err, "db error")
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	_, mock, repo := InitMocks(t)

	// Define the album to be updated
	updatedAlbum := repository.Album{Title: "Updated Album", Artist: "Updated Artist", Price: 10.99}
	idToUpdate := "existing-id"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery("UPDATE \"dev-schema\".albums SET").
			WithArgs(updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, idToUpdate).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "artist", "price"}).AddRow(idToUpdate, updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price))

		updatedAlbumResult, err := repo.Update(idToUpdate, updatedAlbum)

		assert.NoError(t, err)
		assert.NotNil(t, updatedAlbumResult)
		assert.Equal(t, idToUpdate, updatedAlbumResult.ID)
		assert.Equal(t, updatedAlbum.Title, updatedAlbumResult.Title)
		assert.Equal(t, updatedAlbum.Artist, updatedAlbumResult.Artist)
		assert.Equal(t, updatedAlbum.Price, updatedAlbumResult.Price)
	})

	t.Run("Album not found", func(t *testing.T) {
		mock.ExpectQuery("UPDATE \"dev-schema\".albums SET").
			WithArgs(updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, "non-existing-id").
			WillReturnError(sql.ErrNoRows)

		_, err := repo.Update("non-existing-id", updatedAlbum)

		assert.EqualError(t, err, "sql: no rows in result set")
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery("UPDATE \"dev-schema\".albums SET").
			WithArgs(updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, "any-id").
			WillReturnError(errors.New("db error"))

		_, err := repo.Update("any-id", updatedAlbum)

		assert.EqualError(t, err, "db error")
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
