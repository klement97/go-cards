package main

import (
	"os"
	"testing"
)

func TestNewDeckLength(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 52 but got ", len(d))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spaces but got ", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(deck) != len(loadedDeck) {
		t.Error("Expected to have the same length from both decks.")
	}

	os.Remove(filename)
}
