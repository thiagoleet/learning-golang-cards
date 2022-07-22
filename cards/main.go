package main

func main() {
	cards := newDeck()
	cards.suffle()
	cards.print()
}
