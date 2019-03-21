package main

import "fmt"

func main() {
	/*
		//var card string = "Ace of Spades"
		card := "Ace of Spades" // Initialization statemetn
		//card = "Five of Diamonds" // Assigning  new value

		card = newCard()
		fmt.Println(card)

	*/

	/*
		cards := []string{newCard(), "Ace of Diamonds"}
		cards = append(cards, "Six of Spades")
		fmt.Println(cards)

		for i, card := range cards {
			fmt.Println(i, card)
		}
	*/

	/*
		cards := deck{"Ace of Diamonds", newCard()}
		fmt.Println(cards)
		cards.print()
	*/

	/*
		cards := newDeck()
		cards.print()
	*/

	cards := newDeck()
	cards.print()
	fmt.Println(cards[0:9])
	fmt.Println(cards[0:])
	fmt.Println(cards[:9])

}

/*
func newCard() string {
	return "Five of Diamonds"
}
*/
