package main

import (
	"fmt"
)

// eg 2
func equiTriangle(base , height float64) (area , perimeter float64) {
	area = 0.5 * base * height
	perimeter = 3 * base
	return // named return types
}

func main() {
	base, height := 10.5, 4.6
	area, _  := equiTriangle(base, height)
	fmt.Printf("Area is %f\n", area)
}

// eg 1
func calcAmount(price, qty int) int {
	var totalAmount = price * qty
	return totalAmount
}

func main() {
	price, qty := 20, 5
	totalAmount := calcAmount(price, qty)
	fmt.Printf("Total AMount %d\n", totalAmount)
}