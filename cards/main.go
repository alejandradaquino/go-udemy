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
	println(cards.toString())
	cards.save("deck.txt")
	retrieved := newDeckFromFile("deck.txt")
	println(" \nretrieved: \n")
	retrieved.print()
	cards.shuffle()
	println(" \nShuffled:  \n")
	cards.print()
}
