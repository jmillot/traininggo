package application

import "database/sql"

type application struct {
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

func dbConnaect(db *sql.DB) {

}
