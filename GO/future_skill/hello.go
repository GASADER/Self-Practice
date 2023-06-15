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

	//raw string
	data := `1:a 
2:b
3:c`
	fmt.Println(data)

	//template  literal
	result := fmt.Sprintf("have %d THB", price)
	fmt.Printf(result)

	var char rune = 'A'
	fmt.Println("char:", char)

	//Fromatting
	mock := 123

	//string & new line
	fmt.Println("")

	//format string
	fmt.Printf("mock: %d", mock)

	// %s ใช้สำหรับแทนที่ string
	// %d ใช้สำหรับแทนที่ integer
	// %f (float)
	// %.2f(float 2 position)
	// %t (boolean)
	// %v (ค่าประเภทอื่นๆ)
	// %#v all type
	// %c character
	// \n new line
	// %T find type

}
