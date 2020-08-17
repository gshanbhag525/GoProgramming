package main

import "fmt"

func main() {

	//############## Basic For Loop Clause ##############
	sum := 0
	for i := 0 ;i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	//############## optional : init and post statements ##############
	// sum := 0
	// i := 0 
	// for  ;i < 10; {
	// 	sum += i
	// 	i++
	// }
	// fmt.Println(sum)


	//############## while loop ##############
	// sum := 1
	// for sum < 100 {
	// 	sum += sum
	// }
	// fmt.Println(sum)	


	//############## infinite loop ##############
	// sum := 1
	// for {
	// 	fmt.Println(sum)
	// }


	//############## break stmt in infinte loop ##############
	// sum := 1
	// for {
	// 	sum += sum
	// 	fmt.Println(sum)
	// 	if sum == 8 {
	// 		break
	// 	}
	// }


	//############## continue stmt in infinte loop ##############
	// sum := 1
	// for {
	// 	sum += sum
	// 	if sum == 8 {
	// 		continue
	// 	}
	// 	fmt.Println(sum)
	// 	if sum == 32 {
	// 		break
	// 	}
	// }
	
}
