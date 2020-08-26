package application

import (
	"database/sql"
)

// Application is the main object
type Application struct {
	ID             int
	Name           string
	Business       string
	Domaine        int   // Couvre tout ou partie
	Application    int   // Interagit avec
	Technologie    int   // S'appuie sur
	Serveur        int   // Est installé sur
	Bdd            int   // Utilise
	Data           []int // Manipule
	Flux           []int // Emet ou Reçois
	Fonctionnalite []int // Met à disposition
}

func dbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "architect_mtp"
	dbPass := "architect"
	dbName := "architecture_mtp"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
