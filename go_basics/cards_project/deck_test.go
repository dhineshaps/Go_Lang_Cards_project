package main

import (
	"os"
	"testing"
)

// import "testing"

func TestNewDeck(t *testing.T) { //t is the test handler if something goes wrong it will be told

	d := newDeck()

	if len(d) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//test if first card is

	if d[0] != "Ace of Spades" {

		t.Errorf("Expected first of the card is Ace of Spades , but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {

		t.Errorf("Expected last of the card is our of Clubs , but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
