package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionary"

	"training.go/hangman/hangman"
)

func main() {

	err := dictionary.LoadFile("words.txt")
	if err != nil {
		fmt.Println("Probleme de chargement du dictionnaire")
		os.Exit(0)
	}

	g := hangman.New(8, dictionary.PickWord())

	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)

		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("COuld not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
