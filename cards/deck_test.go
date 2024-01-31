package main

import (
	"os"
	"slices"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	wantLen := 24
	if len(d) != wantLen {
		t.Errorf("got length of deck %v want %v", len(d), wantLen)
	}

	wantFirstItem := card{suit: spades.String(), value: "Ace"}

	if d[0] != wantFirstItem {
		t.Errorf("got first item of the deck %v want %v", d[0], wantFirstItem)
	}

	hasKingOfDiamonds := slices.Contains(d, card{suit: diamonds.String(), value: "King"})

	if !hasKingOfDiamonds {
		t.Errorf("deck does not contain King of diamonds")
	}
}

func TestSaveToFileAndNewFromFile(t *testing.T) {
	d := newDeck()
	testfileName := "_decktesting"

	// clean up
	os.Remove(testfileName)

	err := d.saveToFile(testfileName)

	if err != nil {
		t.Errorf("error during save to file call: %v", err)
	}

	loadedDeck := newFromFile(testfileName)

	wantLen := 24
	if len(loadedDeck) != wantLen {
		t.Errorf("got length of deck %v want %v", len(loadedDeck), wantLen)
	}

	// cleanup
	os.Remove(testfileName)

}
