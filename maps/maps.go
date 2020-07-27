package main

import (
	"fmt"
)

func main() {
	countryPopulation := make(map[string]int, 10)

	countryPopulation = map[string]int{

		"India" : 123,
		"USA" : 342,
		"France" : 543,
		"China" : 876,
		"Pakistan" : 678,
	}
	fmt.Println(countryPopulation)
	cp := countryPopulation
	delete(countryPopulation, "USA")
	fmt.Println(cp)
	fmt.Println(countryPopulation)
	
}


