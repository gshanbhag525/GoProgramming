package main

import (  
    "fmt"
)

func accountDetails(userID *int, branchCode *int) {  
    defer fmt.Println("deferred call in accountDetails")

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

