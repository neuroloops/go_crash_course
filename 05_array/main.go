package main

import "fmt"

func main() {
	var fruitArr [2]string

	// assing values
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	//declare and assing

	fruitArr2 := [2]string{"apple", "orange"}

	fruitSlice := []string{"apple", "orange", "grape", "cherry"}

	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr2)
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])

}
