package main

import "fmt"

func main() {
	// initializing slice in go
	cards := []string{"Ace of diamonds", newCard()}
	//appending new thing in slice is append(nameofslice, value you want to append)
	cards = append(cards, "Six of spades")
	//itrate over slice cards slice name we want to itrate
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {

	return "Five Of Diamonds"
}
