package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	// above and below are same
	fmt.Println(card)
}

func newCard() string {
	return "10 of spade"
}
