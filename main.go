package main

func main() {
	cards := newDeck()
	// returns a new slice
	cards = append(cards, "Six of Hearts")
	// range is a keyword for iterating over a slice
	// in a for loop you reinitialize i and card on each iteration
	// so that's why you use := syntax
	cards.print()
}
