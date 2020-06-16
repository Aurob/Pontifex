package main

import "fmt"

func main() {
	cards := newDeck()
	//hand, _ := deal(cards, 5)
	//hand.saveToFile("cards")
	//hand := deckFromFile("cards")

	//cards.shuffle()
	fmt.Println(" ")
	key("Hello", cards)

}

func key(s string, cards deck) {
	jokerAPos := cards.findValue("Joker A")
	cards.moveDown(jokerAPos, 1)
	jokerBPos := cards.findValue("Joker B")
	cards.moveDown(jokerBPos, 2)

	//cards.print()
	fmt.Println(" ")

	c3 := cards.tripleCut()
	//c3.print()

	cc := c3.countCut()
	//cc.print()

	outputCard := cc[cc[0].value]
	fmt.Println(outputCard.value)
	cards = cc
}
