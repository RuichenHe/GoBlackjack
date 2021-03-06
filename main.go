package main

import (
	"fmt"

	"github.com/RuichenHe/GoBlackjack/blackjack"
)

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)

}
