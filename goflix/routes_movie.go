package main

import (
	"log"
	"net/http"
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

func mapMovieToJSON(m Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
