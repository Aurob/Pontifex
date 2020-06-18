package main

import (
	"fmt"
)

var alpha string = "abcdefghijklmonpqrstuvwxyz"

func main() {
	cards := newDeck()
	//cards.shuffle()
	fmt.Println(key("1234567890", cards))
}

func key(s string, cards deck) (deck, string) {
	var keystream string
	for range s {
		jokerAPos := cards.findValue("Joker A")
		cards.moveUp(jokerAPos, 1)
		jokerBPos := cards.findValue("Joker B")
		cards.moveUp(jokerBPos, 2)

		cards = cards.tripleCut()
		cards = cards.countCut()

		outputCard := cards.getOutput()
		if outputCard.value == 53 {
			var ks string
			cards, ks = key("1", cards)
			keystream += ks
		} else {
			keystream += string(alpha[outputCard.value%26])
		}
	}
	return cards, keystream
}
