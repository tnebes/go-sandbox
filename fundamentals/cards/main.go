package main

import "fmt"

func main() {
	var cards = []string{newCard(), newCard()}

	fmt.Println(cards)
	// one way of doing it
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// alternative way of doing it
	for i := 0 ; i < len(cards) ; i++ {
		fmt.Println(cards[i])
	}

}

func newCard() string {
	return "Five of Diamonds"
}
