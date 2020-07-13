package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLength := 40
	firstCard := "Ace of Spades"
	lastCard := "Ten of Clubs"

	if len(d) != deckLength {
		t.Error("Expecting deck length of ", deckLength, ", but got ", len(d))
	}

	if d[0] != firstCard {
		t.Error("The first card must be ", firstCard, ", but got ", d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Error("The last card must be ", lastCard, ", but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	d := newDeck()

	fileName := "_decktesting"
	// Clean up before testing
	os.Remove(fileName)

	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	deckLength := 40
	firstCard := "Ace of Spades"
	lastCard := "Ten of Clubs"

	if len(loadedDeck) != deckLength {
		t.Error("Expecting deck length of ", deckLength, ", but got ", len(loadedDeck))
	}

	if loadedDeck[0] != firstCard {
		t.Error("The first card must be ", firstCard, ", but got ", loadedDeck[0])
	}

	if loadedDeck[len(d)-1] != lastCard {
		t.Error("The last card must be ", lastCard, ", but got ", loadedDeck[len(d)-1])
	}

	// Clean up after testing
	os.Remove(fileName)
}
