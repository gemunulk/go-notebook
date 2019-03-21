package main

import "fmt"

// Create a new typ of 'deck'
// which is a slice of strings

type deck []string // in OOP ~= deck is a class

// in OOP ~=  print is a method under deck class
// (d deck) is called 'reciever'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
func (this deck) print() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}
*/

/*
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
*/

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
