# 🃏 standard-deck

A library to create, shuffle, and play with a Standard Deck of Playing Cards. Allows you to focus on the game's rules
while this library manages the deck itself.

# Installation

```shell
go get github.com/bytebury/standard-deck
```

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

    // Get the value and display value of your first card
    fmt.Printf("You have a %s of %s\n", myHand[0].DisplayValue, myHand[0].Suit)

    // For weighted values of a card, you can get the `Value`
    // Ace has weighted value of 1 and King has weighted value of 13
    fmt.Println("Card value is:", myHand[0].Value)
}
```
