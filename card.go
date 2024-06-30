package standarddeck

import (
	"fmt"

	"github.com/bytebury/standard-deck/suit"
)

type Card struct {
	Suit         suit.Suit
	DisplayValue string
	Value        int
}

func newCard(suitValue int, value int) Card {
	return Card{
		Suit:         getSuitByValue(suitValue),
		DisplayValue: getDisplayValue(value),
		Value:        value,
	}
}

func getSuitByValue(value int) suit.Suit {
	switch value {
	case 0:
		return "Heart"
	case 1:
		return "Diamond"
	case 2:
		return "Club"
	case 3:
		return "Spade"
	}
	panic("Unable to get suit by the given value")
}

func getDisplayValue(value int) string {
	switch value {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return fmt.Sprintf("%d", value)
	}
}
