package main

import "fmt"
import "strconv"

func main() {

	var a int = 12
	fmt.Printf("%v, %T", a, a)

	var b string
	b = strconv.Itoa(a)
	fmt.Printf(" %v, %T", b, b)


}
