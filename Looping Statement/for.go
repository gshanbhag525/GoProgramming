package main

import "fmt"

func main() {

	//############## Basic For Loop Clause ##############
	// sum := 0
	// for i := 0 ;i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)


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
	
	
	//############## For loop with range clause ##############
	
	// slices
	
	nums := []string{"one", "two", "three"	}

	for idx, num := range nums {
		fmt.Printf("%d : %s\n", idx, num)
	}

	for idx := range nums {
		fmt.Printf("%d \n", idx)
	}

	
	for _, num := range nums {
		fmt.Printf("%s \n", num)
	}

	// array

	nums := [...]string{"one", "two", "three"	}

	for idx, num := range nums {
		fmt.Printf("%d : %s\n", idx, num)
	}

	for idx := range nums {
		fmt.Printf("%d \n", idx)
	}

	for _, num := range nums {
		fmt.Printf("%s \n", num)
	}

	for range nums {
		fmt.Printf("gunesh")
	}


	// pointer to array
	
	nums := &[...]string{"one", "two", "three"	}

	for idx, num := range *nums {
		fmt.Printf("%d : %s\n", idx, num)
	}

	for idx, num := range nums {
		fmt.Printf("%d : %s\n", idx, num)
	}

	// maps

	m := map[string]int{"one" : 1, "two": 2, "three": 3	}

	for k := range m {
		fmt.Printf("keys : %q\n", k)
	}

	for k, v := range m {
		fmt.Printf("keys : %q, value : %d\n", k, v)
	}
	
	// short variable declaration and scoping

	// idx variable declared outside for loop
	m := []string{"zero","one", "two"}
	var idx int
	for idx = range m {
	}
	fmt.Println(idx)

	// idx variable value not accessible outside the for loop
	m := []string{"zero","one", "two"}
	for idx := range m {
	}
	fmt.Println(idx)



}

