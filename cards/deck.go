package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) save() {
	err := ioutil.WriteFile("deck.txt", []byte(d.toString()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
