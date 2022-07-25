package blackjack

import (
	"fmt"

	deck "github.com/RuichenHe/GoCard-module"
)

type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}
func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {

}
func HumanAI() AI {
	return humanAI{}
}

type humanAI struct {
}

func (ai humanAI) Bet(shuffled bool) int {
	fmt.Println("What would you like to bet??")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}
func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		var input string
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it? (s)stand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL CARDS==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}
