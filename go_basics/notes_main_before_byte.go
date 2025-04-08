package main

import "fmt"

func main() {

	cards := newDeck()

	//cards.print()

	hand, remainingDeck := deal(cards, 5) //calling func deal with args
	//which has two return variable s given so two retrun O/P is given

	hand.print()
	fmt.Println("handful cards selected")
	remainingDeck.print()
