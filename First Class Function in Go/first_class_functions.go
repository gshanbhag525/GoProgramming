package main

import (
	"fmt"
)

func main() {
	// Assigning variable to Anonymous function
	first := func() {
		fmt.Println("Inside the Anonymous function")
	}
	first()
	fmt.Printf("%T", first)


	// Anonymous function
	func() {
		fmt.Println("Inside the Anonymous function")
	}()
	
	// Passing String to Anonymous function
	func(str string) {
		fmt.Println("Inside the Anonymous function", str)
	}("Hello There")
}

// User Defined Function types
package main

import (
	"fmt"
)

type add func(a int, b int) int

func main() {
	var a add = func(a , b int) int {
		return a+b
	}
	s := a(50, 5)
	fmt.Println("sum ", s)
}


// Passing functions as arguments
package main

import (
	"fmt"
)

func simple(a func(a,b int) int) {
	fmt.Println(a(50, 5))
}

func main() {
	f := func(a, b int) int {
		return a+b
	}
	simple(f)
}


// returning functions from other functions
package main

import (
	"fmt"
)

func simple() func(a,b int) int {
	f := func(a, b int) int {
		return a+b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(50, 5))
}