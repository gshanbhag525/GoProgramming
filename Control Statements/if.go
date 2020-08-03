package main

import (
	"fmt"
)

func main() {

	//  Relational Operator ( < > <= >= ==)
	
	a := 10
	b := 20
	if a > b {
		fmt.Println("A is bigger ")
	}
	if a < b {
		fmt.Println("B is bigger ")
	}
	if a == b {
		fmt.Println("Both are equal")
	}

	// a := 150
	// b := 40
	// c := 50

	// Logical Operators ( && || !)

	// if a > b && a > c {
	// 	fmt.Println("A is biggest ")
	// } else if b > a && b > c {
	// 	fmt.Println("B is biggest ")
	// } else {
	// 	fmt.Println("C is biggest ")
	// }

	// Or operator usage

	// if a > 50 || a < 100 {
	// 	fmt.Println("A lies in between 50 , 100")
	// }

	// Not operator usage

	if !false {
		fmt.Println("True")
	}
	
	// Error handling in Go
	
	// if err != nil {
	// 	// handle the error
	// }

	// Initializing in if statement

	if num := 100; num > 0 {
		fmt.Println("num is Positive")
	}

	if num := 100; num%2 == 0 {
		fmt.Println("num is Even")
	}



	
}