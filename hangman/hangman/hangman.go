package hangman

import "strings"

// Game is the struct of the game
type Game struct {
	State        string   // Game state
	Letters      []string // Letters in the word to find
	FoundLetters []string // Good guesses
	UsedLetters  []string // Used letters
	TurnsLeft    int      // Remaiing attemps
}

// New return a pointer on a new Game struct
func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return g
}

// MakeAGuess for status managment
func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWorld(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWorld(guess, g.Letters) {
		g.State = "goddGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.LooseTurn(guess)
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

// RevealLetter for letter revelation
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

// LooseTurn for wrong letter
func (g *Game) LooseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func letterInWorld(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
