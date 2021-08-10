package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards, hand := deal(cards, 5)
	hand.print()
	cards.print()
}
