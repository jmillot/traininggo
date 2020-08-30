package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStore struct {
	movieID int64
	movies  []*Movie
}

func (t testStore) Open() error {
	return nil
}

func (t testStore) Close() error {
	return nil
}

func (t testStore) GetMovies() ([]*Movie, error) {
	return t.movies, nil
}

func (t testStore) GetMovieByID(id int64) (*Movie, error) {
	for _, m := range t.movies {
		if m.ID == id {
			return m, nil
		}
	}
}

func (t testStore) CreateMovie(m *Movie) error {
	t.movieID++
	m.ID = t.movieID
	t.movies = append(t.movies, m)
	return nil
}
func TestCreateMovieUnit(t *testing.T) {

	log.Print("Démarrage du server")
	srv := newServer()
	srv.store = &testStore{}

	p := struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"relese_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}{
		Title:       "titre du film",
		ReleaseDate: "2020-08-30",
		Duration:    120,
		TrailerURL:  "url de test",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	assert.Nil(t, err)

	log.Printf("Serveur HTTP de test")
	r := httptest.NewRequest("POST", "api/movies/", &buf)
	w := httptest.NewRecorder()

	log.Printf("Traitement de la requete")
	srv.handleCreateMovie()(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
