package main

import (
	"fmt"

	"training.go/Struct/player"
)

// Personne is a type of personne
type Personne struct {
	Nom    string
	Prenom string
}

func main() {
	var pers Personne
	var p1 player.Player

	p1.Name = "Jerome"

	pers.Nom = "Millot"
	pers.Prenom = "Jerome"
	fmt.Printf("Persone: %v\n", pers)

	a := player.Avatar{"http://www.google.com/"}
	fmt.Printf("a: %v\n", a)

	p3 := player.Player{
		Name: "John",
		Age:  25,
		Avatar: player.Avatar{
			URL: "https://www.google.com",
		},
	}
	fmt.Printf("P3: %v", p3)
}
