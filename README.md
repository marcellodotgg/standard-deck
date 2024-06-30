# 🃏 Deck

A library to create, shuffle and play with a Standard Deck of Playing Cards.

# Usage

```go
package main

import (
	"fmt"

	standarddeck "github.com/bytebury/standard-deck"
	"github.com/bytebury/standard-deck/suit"
)

func main() {
	// Create a deck of cards and shuffle them.
	currentDeck := standarddeck.New().Shuffle()

	// Count the number of cards
	fmt.Println("Number of cards:", len(currentDeck.Cards)) // 52

	// Draw 7 cards from the deck
	myHand, err := currentDeck.Draw(7)

	// You'll need to check before you draw, otherwise, you'll get an error if you
	// draw more cards than are available.
	if err != nil {
		fmt.Printf("There are only %d cards in the deck left\n", len(currentDeck.Cards))
	}

	// 7 cards from the start of the deck
	fmt.Println(myHand)

	// Check to see if your first card is a Heart
	fmt.Println(myHand[0].Suit == suit.Heart)
}
```