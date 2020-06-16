package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected \"Ace of Spades\" for first card, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected \"King of Clubs\" for last card, got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndDeckFromFile(t *testing.T) {
	os.Remove(".test_deck")
	deck := newDeck()
	deck.saveToFile(".test_deck")

	loadedDeck := deckFromFile(".test_deck")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(loadedDeck))
	}

	os.Remove(".test_deck")
}
