package main

import "fmt"

func main() {

	nos := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(nos)

	for _, val := range nos {

		//fmt.Println(val % 2)

		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}

	}

}
