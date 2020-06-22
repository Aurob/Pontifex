package main

var alpha string = "abcdefghijklmonpqrstuvwxyz"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.generateKeystream(100)
	cards.print()
}
