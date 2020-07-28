// package main

// import "fmt"

// // type Celebrity struct {
// // 	Age int
// // 	CelebrityName string
// // 	Movies []string
// // 	CoActors []string
// // }

// func main(){

// 	aCelebrity := struct{age int} {age : 55}
// 	bCelebrity := &aCelebrity
// 	bCelebrity.age = 95 
// 	fmt.Println(aCelebrity)
// 	fmt.Println(bCelebrity)

// }


//################## Composition in Golang ########

package main

import "fmt"

type Vehicle struct {
	Name string
	Type string
}

type Car struct {
	Vehicle
	Maxspeed float32
	Fueltype string
}

func main() {
	c := Car{
		Vehicle : Vehicle{Name : "Ferrari", Type : "Convertible" },
		Maxspeed : 250,
		Fueltype : "Diesel",
	}

	fmt.Println(c)

}