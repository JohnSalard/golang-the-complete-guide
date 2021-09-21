package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	d.printDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	// os.Remove(filename)
	deck := newDeck()
	deck.saveToFile(filename)

	loadDeck := newDeckFromFile(filename)
	if len(loadDeck) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(loadDeck))
	}
	fmt.Println(loadDeck)
	os.Remove(filename)
}
