package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()

	hand := (&cards).deal(5)

	hand.print()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
