package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	t.Parallel()
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expect deck length of 52, but got %v", len(d))
	}

	const expectedFirstCard = "Ace of Spades ♠"

	if d[0] != expectedFirstCard {
		t.Errorf("Expect first card of %v, but got %v", expectedFirstCard, d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	t.Parallel()

	const filename = "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expect deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
