package repository_tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodrigoenzohernandez/web-service-gin/internal/repository"
	"github.com/stretchr/testify/assert"
)

func NewAlbumRepo(db *sql.DB) *repository.AlbumRepository {
	return &repository.AlbumRepository{DB: db}
}

func InitMocks(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *repository.AlbumRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	t.Cleanup(func() { db.Close() })

	repo := NewAlbumRepo(db)
	return db, mock, repo
}

type AlbumRepository struct {
	DB *sql.DB
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
		t.Errorf("there were unfulfilled expectations: %s", err)
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
		t.Errorf("there were unfulfilled expectations: %s", err)
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
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
