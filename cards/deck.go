package main

import "fmt"

//Create a new type deck
type deck []string

func newDeck() deck {

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	var hand = d[:handSize]
	var rest = d[handSize:]
	return hand, rest
}
