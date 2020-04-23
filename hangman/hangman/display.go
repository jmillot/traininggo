package hangman

import "fmt"

// DrawWelcome message d'accueil
func DrawWelcome() {
	fmt.Println(`
	#     #    #    #     #  #####  #     #    #    #     # 
	#     #   # #   ##    # #     # ##   ##   # #   ##    # 
	#     #  #   #  # #   # #       # # # #  #   #  # #   # 
	####### #     # #  #  # #  #### #  #  # #     # #  #  # 
	#     # ####### #   # # #     # #     # ####### #   # # 
	#     # #     # #    ## #     # #     # #     # #    ## 
	#     # #     # #     #  #####  #     # #     # #     #
	`)
}

// Draw the game
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
 _____
 |    |
 |    :p
 |   /-\
 |    |
 |   / \
_|_______
	`
	case 1:
		draw = `
 _____
 |    |
 |    O
 |   /|\
 |    |
 |   / \
_|_______
	`
	case 2:
		draw = `
 _____
 |    |
 |    O
 |   /|
 |    |
 |   / \
_|_______
	`
	case 3:
		draw = `
 _____
 |    |
 |    O
 |    |
 |    |
 |   / \
_|_______
	`
	case 4:
		draw = `
 _____
 |    |
 |    O
 |    |
 |    |
 |   /
_|_______
	`
	case 5:
		draw = `
 _____
 |    |
 |    O
 |    |
 |    |
 |
_|_______
	`
	case 6:
		draw = `
 _____
 |    |
 |    O
 |    |
 |    
 |
_|_______
	`
	case 7:
		draw = `
 _____
 |    |
 |    O
 |
 |
 |
_|_______
	`
	case 8:
		draw = `
 _____
 |    |
 |
 |
 |
 |
_|_______
	`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You loose : The word was :")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You win ! the word was :")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
