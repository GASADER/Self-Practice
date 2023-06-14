package main

import (
	"fmt"
)

func main() {
	// variable out scoop
	// var name string = "hello"

	//variable in scoop function
	name := "hello"
	fmt.Println(name)

	//mutiple declaration
	price, age, check := 200, 20, true
	fmt.Println(price, age, check)

	//template  literal
	result := fmt.Sprintf("have %d THB", price)

	fmt.Printf(result)
}
