package main

import "fmt"

type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

func (m Movie) ToString() string {
	return fmt.Sprintf("id=%v, Title=%v, Realease date=%v, Duration=%v, Trailer URL=%v", m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
}
