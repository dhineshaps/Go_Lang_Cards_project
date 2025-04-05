package main

func main() {

	// var card string = "Ace of strings"
	//card := "Ace of strings" //only during the initialization we have to use := for next stime we can use "="
	//card = "new value"
	//card := newCard //to get return from the function
	//fmt.Println(card)
	//fmt.Println(newCard()) //VALID

	//Slice   - is  a list? in python
	cards := deck{"Ace of diamond", newCard()} //here is deck is the type which we have created - in deck.go which inherits the type string
	cards = append(cards, "Six of Spades")

	// //slice iteration , i is index number
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)

	cards.print()
}

// bool - True pr False
// string - "Hi" or " Hows it is going"
//int - 0 , -10000,99999
//float64 - 10.00001, 0.00009, -100.003

// func newCard() string { //need to specify the type of datatype it will return

// 	return "Five of diamonds"
// }
