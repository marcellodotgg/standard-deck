package standarddeck

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	deck := New()

	if len(deck.Cards) != 52 {
		t.Error("Should have 52 cards in the deck")
	}
}

func Test_Shuffle(t *testing.T) {
	deck := New().Shuffle()

	if reflect.DeepEqual(deck.Cards, New().Cards) {
		t.Error("Should have shuffled the deck")
	}
}

func Test_ShuffleInPlace(t *testing.T) {
	deck := New()
	deck.Shuffle()

	if reflect.DeepEqual(deck.Cards, New().Cards) {
		t.Error("Should have shuffled the deck in place")
	}
}

func Test_Draw_TooManyReturnsError(t *testing.T) {
	deck := New()

	if _, err := deck.Draw(53); err == nil {
		t.Error("Should have thrown not enough cards in the deck error")
	}
}

func Test_Draw_LessThanOne(t *testing.T) {
	deck := New()

	if _, err := deck.Draw(0); err == nil {
		t.Error("Should have thrown too cannot draw less than 1 card")
	}
}

func Test_Draw_RemovesFromDeck(t *testing.T) {
	deck := New()

	cards, _ := deck.Draw(7)

	if len(cards) != 7 {
		t.Errorf("Expected to draw 7 cards, actually drew %d", len(cards))
	}

	if len(deck.Cards) != 45 {
		t.Errorf("Expected deck to have 45 cards, actually had %d", len(deck.Cards))
	}
}
