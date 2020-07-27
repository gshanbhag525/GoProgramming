package main

import "fmt"

// type Celebrity struct {
// 	Age int
// 	CelebrityName string
// 	Movies []string
// 	CoActors []string
// }

func main(){

	aCelebrity := struct{age int} {age : 55}
	bCelebrity := &aCelebrity
	bCelebrity.age = 95 
	fmt.Println(aCelebrity)
	fmt.Println(bCelebrity)

}