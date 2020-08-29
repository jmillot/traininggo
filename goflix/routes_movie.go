package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

func (s *server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies()
		if err != nil {
			log.Printf("Cannot load movies, err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var reponse = make([]jsonMovie, len(movies))
		for i, m := range movies {
			reponse[i] = mapMovieToJSON(m)
		}
		s.respond(w, r, reponse, http.StatusOK)
		return
	}
}

func (s *server) handleMovieDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vars = mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Invalid parameter, err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		movie, err := s.store.GetMovieByID(id)
		if err != nil {
			log.Printf("Cannot load Movie, err=%v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var reponse jsonMovie
		reponse = mapMovieToJSON(movie)
		s.respond(w, r, reponse, http.StatusOK)
		return
	}
}

func mapMovieToJSON(m *Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}