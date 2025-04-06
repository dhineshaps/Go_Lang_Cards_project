package main

import "testing"

// import "testing"

func TestNewDeck(t *testing.T) { //t is the test handler if something goes wrong it will be told

	d := newDeck()

	if len(d) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

}
