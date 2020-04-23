package hangman

import "testing"

func TestLetterInWorld(t *testing.T) {
	word := []string{"h", "e", "l", "l", "o"}
	guess := "o"
	hasLetter := letterInWorld(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}
