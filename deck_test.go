package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	// First card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	// Last card is King of Diamonds
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected King of Diamonds but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Invalid deck length. Expected 52, found %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
