package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Store interface for database management
type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
	GetMovieByID(id int64) (*Movie, error)
	CreateMovie(m *Movie) error
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	db.MustExec(schema)
	store.db = db
	return nil

}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (store *dbStore) GetMovieByID(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		log.Printf("Erreur dans store.go method GetMovieByID")
		return nil, err
	}
	return movie, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {
	result, err := store.db.NamedExec(`INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (:title,:release,:duration,:url)`,
		map[string]interface{}{
			"title":    m.Title,
			"release":  m.ReleaseDate,
			"duration": m.Duration,
			"url":      m.TrailerURL,
		})
	if err != nil {
		return err
	}
	m.ID, err = result.LastInsertId()
	return err
}
