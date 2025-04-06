package main

import (
	"fmt"
	"os"
	"strings"
)

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

//helper function to convert the deck to string to store in the file

func (d deck) toString() string { //reciever function
	return strings.Join([]string(d), ",")

}

//to save to file

func (d deck) saveToFile(filename string) error { //if there is any error it will return error

	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

//to ead the file

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //byteSlice err := os.ReadFile(filename) -- err type will have nil if nothing went wrong

	if err != nil {
		//option #1 - log the error and return a call to newDeck()
		//option #2 - log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") //converting byte to string and split by "," to slice of string

	return deck(s) //coverting slice of string into deck
}
