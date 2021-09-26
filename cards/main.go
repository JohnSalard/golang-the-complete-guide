package main

func main() {

	// cards := newDeck()
	// cards.print()

	// cards := newDeck()
	// hand, remain := deal(cards, 5)
	// hand.print()
	// remain.print()
	// greeting := "สวัสดี"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.saveToFile("ex.txt")
	result := newDeckFromFile("ex.txt")
	result.print()

	cards.shuffle()
	// cards.printDeck()
	// fmt.Println(cards)

}
