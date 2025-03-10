package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 3)

	// hand.print()

	// remainingDeck.print()

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}
