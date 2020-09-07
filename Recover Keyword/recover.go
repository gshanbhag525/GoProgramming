package main

import (  
	"fmt" 	
	"runtime/debug"
)

func recoverAccount() {
	if r := recover(); r != nil {
		fmt.Println("recovered from \n", r)
		debug.PrintStack()
	}
}

func accountDetails(userID *int, branchCode *int) {  
    defer recoverAccount()

    if userID == nil {
        panic("runtime error : user id cant be nil")
    }

    if branchCode == nil {
        panic("runtime error : branch code cant be nil")
    }
    fmt.Printf("%d %d\n", *userID, *branchCode)
    fmt.Println("returned normally from accountDetails function")
}

func main() {  
    defer fmt.Println("deferred call in main function")
    userID := 4567
    // branchCode := 12
    accountDetails(&userID, nil)
    fmt.Println("returned normally from main function")
}

// eg 2

package main

import (  
    "fmt"
)

func recoverSlicePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func slicePanic() {  
	defer recoverSlicePanic()
    n := []int{5, 7, 4}
    fmt.Println(n[4])
    fmt.Println("normally returned from a")
}
func main() {  
    slicePanic()
    fmt.Println("normally returned from main")
}