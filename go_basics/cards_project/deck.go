package main

import "fmt"

// create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck { //this function returns list of strings as cards , deck is type string in slice

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"} //string values
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() { //reciever function  //here now d is cards slice variable which it got from deck //it can be aby char or word

	//slice iteration , i is index number
	for i, card := range d { //for i , card := range cards

		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { //here function deal has two arguments, so it needs to be called by two args
	//it returns two values of type deck

	return d[:handSize], d[handSize:]

}
