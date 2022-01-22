package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create new type of 'deck'
// that is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			card := suit + " of " + value
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
// d of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}