package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	fmt.Println("Head 1")
	head, remainingCards := deal(cards, 3)
	head.print()
	fmt.Println("Head 2")
	head2, _ := deal(remainingCards, 3)
	head2.print()
} 

func newCard() string {
	return "Five of Diamonds"
}