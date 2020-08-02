package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string `maxsize:"20"`
	Password string `minsize:"8"`
}

func main() {
	u := User{}
	t := reflect.TypeOf(u)
	// field := t.Field(0) // alternative
	field, _ := t.FieldByName("Password")
	fmt.Println(field.Tag)
	// fmt.Println(field.Tag.Get("maxsize"))

}