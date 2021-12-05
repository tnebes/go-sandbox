package main

func main() {
	cards := deck{newCard(), newCard(), "Bluh of derp"}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
