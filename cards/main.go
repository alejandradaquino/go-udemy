package main

import "fmt"

var test int

func main() {

	test = 9

	otherCard := "Saraza"

	var cards = []string{initialCard(), otherCard}
	cards = append(cards, "other thing")
	cards = append(cards, "missing java :(")
	fmt.Println(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}

func initialCard() string {
	return "Ace of Diamond"
}
