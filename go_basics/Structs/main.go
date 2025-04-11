package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct { //defining struct called "person"

	firstName string
	lastName  string
	contact   contactInfo //embedding struct
}

func main() {

	//Dhinesh := person{"Dhinesh", "Palanisamy"} //declaring the struct - here it automa. takes first val is firs name
	//Dhinesh := person{firstName: "Dhinesh", lastName: "Palanisamy"} //declaring the struct - best way

	// var Dhinesh person

	// Dhinesh.firstName = "Dhinesh" //updating the structs
	// Dhinesh.lastName = "Palanisamy"

	// fmt.Println(Dhinesh)
	// fmt.Printf("%+v", Dhinesh)

	Dhinesh := person{
		firstName: "Dhinesh",
		lastName:  "Palanisamy",
		contact: contactInfo{
			email:   "dhineshpazanisamy@gmail.com",
			zipCode: 636451, //even last line should have commad
		}, //comma should be here after embedded
	}

	fmt.Println(Dhinesh)
	//fmt.Printf("%+v", Dhinesh)
	Dhinesh.print() //invoking the reciever function
	Dhinesh.updateName("Dhine")
	Dhinesh.print()
}
func (p person) updateName(newFirstName string) {

	p.firstName = newFirstName

}

func (p person) print() { //reciever function, here this func can be invoked with type peron
	// and those values will be assigned to p

	fmt.Printf("%+v", p)

}
