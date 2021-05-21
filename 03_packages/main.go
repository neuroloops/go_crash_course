package main

import (
	"fmt"
	"go_crash_course/simpleinterest"
	"math"
)

func main() {

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
