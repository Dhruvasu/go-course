package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// append returns new slice
	cards = append(cards, "Four of Hearts")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
