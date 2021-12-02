package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my")
	// cards.print()
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// cards.print()
	// for index, item := range cards {
	// 	fmt.Println(index, item)
	// }

	// fmt.Println(cards)
}
