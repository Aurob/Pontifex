package main

import "fmt"

var alpha string = "abcdefghijklmonpqrstuvwxyz"

func main() {
	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards.generateKeystream(100))

}
