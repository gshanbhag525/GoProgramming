package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	switch i {
	case 1 : 
		fmt.Println("One")
	case 2 : 
		fmt.Println("Two")
	default :
		fmt.Println("Default")
	}


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: 
		fmt.Println("Its  a Weekend")
		
	default :
		fmt.Println("Its not a Weekend")
		// break
		fmt.Println("Hello welcome")
	}
	
	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("Its before noon")
	default :
		fmt.Println("Its after noon")
	}
	
}

