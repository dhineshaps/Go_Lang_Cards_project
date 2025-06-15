package main

import "fmt"

func main() {

	colors := map[string]string{ //declaring colour with both key and value is string
		"red":   "#ff0000",
		"green": "#4bf45",
		"white": "dfrrgg",
	}

	//another way to decalar maps

	var colors1 map[string]string //decalaring a empty map , which later can be value assigned

	colors2 := make(map[int]string) //decalaring a empty map , which later can be value assigned , key is int and value is string

	colors2[10] = "#gkmfng"

	fmt.Println(colors)
	fmt.Println(colors["green"]) //to access particular key
	fmt.Println(colors1)
	fmt.Println(colors2)

	printMap(colors) //calling func to iterate over

	delete(colors, "green")
	fmt.Println("After deleting")
	fmt.Println(colors)

}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
