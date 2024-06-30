package standarddeck

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

// Shuffles the deck of cards in place and returns it.
func (d Deck) Shuffle() Deck {
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(d.Cards)

	for i := length - 1; i > 0; i-- {
		j := randomNumber.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

	return d
}

// Draws the given number of cards from the deck and returns the cards taken out.
// This will return an error if there are not enough cards in the deck.
func (d *Deck) Draw(number int) ([]Card, error) {
	if number > len(d.Cards) {
		return []Card{}, errors.New("not enough cards in the deck")
	}

	if number <= 0 {
		return []Card{}, errors.New("cannot draw less than 1 card")
	}

	cards := d.Cards[:number]
	d.Cards = d.Cards[number:]

	return cards, nil
}

// Generates a new deck of cards unshuffled.
// The order of a new deck of cards is: Hearts A -> K, Diamonds A -> K, Clubs A -> K, Spades A -> K
func New() Deck {
	var cards []Card

	for suitVal := 0; suitVal < 4; suitVal++ {
		for value := 1; value < 14; value++ {
			cards = append(cards, newCard(suitVal, value))
		}
	}

	return Deck{Cards: cards}
}
