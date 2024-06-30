package standarddeck

import (
	"strconv"
	"testing"

	"github.com/bytebury/standard-deck/suit"
)

func TestNewCard_AceOfHearts(t *testing.T) {
	card := newCard(0, 1)

	if card.Suit != suit.Heart {
		t.Errorf("Suit was incorrect, got %s, wanted %s", card.Suit, suit.Heart)
	}

	if card.DisplayValue != "A" {
		t.Errorf("Value was incorrect, got %s, wanted %s", card.DisplayValue, "A")
	}
}

func TestNewCard_SixOfDiamonds(t *testing.T) {
	card := newCard(1, 6)

	if card.Suit != suit.Diamond {
		t.Errorf("Suit was incorrect, got %s, wanted %s", card.Suit, suit.Diamond)
	}

	if card.DisplayValue != "6" {
		t.Errorf("Value was incorrect, got %s, wanted %s", card.DisplayValue, "6")
	}
}

func TestGetSuitByValue_NoPanic(t *testing.T) {
	heart := getSuitByValue(0)
	diamond := getSuitByValue(1)
	club := getSuitByValue(2)
	spade := getSuitByValue(3)

	if heart != suit.Heart {
		t.Errorf("Suit was incorrect, expected Heart")
	}

	if diamond != suit.Diamond {
		t.Errorf("Suit was incorrect, expected Diamond")
	}

	if club != suit.Club {
		t.Errorf("Suit was incorrect, expected Club")
	}

	if spade != suit.Spade {
		t.Errorf("Suit was incorrect, expected Spade")
	}
}

func TestGetSuitByValue_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function did not panic")
		}
	}()
	// should panic
	getSuitByValue(4)
}

func TestGetDisplayValue_Pip(t *testing.T) {
	// pip cards are ones with faces
	if getDisplayValue(11) != "J" {
		t.Error("Display value should have been J")
	}

	if getDisplayValue(12) != "Q" {
		t.Error("Display value should have been Q")
	}

	if getDisplayValue(13) != "K" {
		t.Error("Display value should have been K")
	}
}

func TestGetDisplayValue_Ace(t *testing.T) {
	if getDisplayValue(1) != "A" {
		t.Error("Display value should have been A")
	}
}

func TestGetDisplayValue_NumberCard(t *testing.T) {
	// numbers start at 2 and end at 10
	for i := 2; i < 11; i++ {
		if getDisplayValue(i) != strconv.Itoa(i) {
			t.Errorf("Display value should have been %s", strconv.Itoa(i))
		}
	}
}
