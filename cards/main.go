package main

func main() {

	var cards = newDeck()
	cards = append(cards, "other thing")
	cards = append(cards, "missing java :( ")
	cards.print()

	deal, rest := cards.deal(3)

	println("Deal: ")
	deal.print()
	println("Rest: ")
	rest.print()
}
