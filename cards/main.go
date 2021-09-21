package main

func main() {

	cards := newDeck()

	// cards.printDeck()

	// hand, remain := deal(cards, 5)
	// hand.printDeck()
	// fmt.Println(hand)
	// fmt.Println(remain)

	// greeting := "Hi there"
	// fmt.Println(greeting)
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())
	// cards.saveToFile("ex.txt")
	// result := newDeckFromFile("ex.txt")
	// fmt.Println(result)

	cards.shuffle()
	// cards.printDeck()
	// fmt.Println(cards)

}
