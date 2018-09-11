package main

func main() {
	// cards := newDeck()
	// cards.show()
	// hand, remainingCard := deal(cards, 5)

	// hand.show()
	// remainingCard.show()
	// // greetings := "Hi there"
	// // fmt.Println([]byte(greetings))
	// cards := newDeck()
	// cards.saveToFile("my cardsc")
	// cards := newDeckFromFile("my cards")
	// cards.show()
	cards := newDeck()
	cards.shuffle()
	cards.show()
}
