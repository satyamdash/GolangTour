package main

import (
	"os"
	"testing"
)

func TestDeck(t *testing.T) {

	d := createDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

}

func TestFirstCard(t *testing.T) {
	d := createDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := createDeck()
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestFileIO(t *testing.T) {
	os.Remove("_decktesting")
	deck := createDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
