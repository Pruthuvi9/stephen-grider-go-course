package main

func main() {
	// card := "Ace of Spades"
	// card := newCard()
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// for i := range cards {
	// 	fmt.Println(cards[i])
	// }

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
