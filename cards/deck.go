package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
	for _, card := range d {
		fmt.Println(card)
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

func (d deck) save(path string) {
	err := ioutil.WriteFile(path, []byte(d.toString()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(path string) deck {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := string(bytes)
	cards := strings.Split(s, ",")
	return deck(cards)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
