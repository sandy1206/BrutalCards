package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	hand, remainingCards := cards.deal(5)
	fmt.Println(hand, "------------------------------------", remainingCards)
}
