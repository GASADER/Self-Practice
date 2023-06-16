package main

import (
	"fmt"
	"math"
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
	var mock rune = '😊'

	//string & new line
	fmt.Println("mock: ", mock)

	//format string
	fmt.Printf("mock: %c \n", mock)

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

	//Zero value

	//nil = null
	var num int
	var digi float32
	var status bool
	var word string
	var emo rune

	fmt.Printf("boo: %t \n", status)  //false
	fmt.Printf("int: %d \n", num)     //0
	fmt.Printf("float: %f \n", digi)  //0.00000
	fmt.Printf("string: %s \n", word) //""
	fmt.Printf("char: %c \n", emo)    //0

	//if else
	num = 34
	if num == 34 && (num > 36 || num < 39) {
		fmt.Printf("Yes: %d \n", num)
	}
	//short if
	w := 25.0
	if v := math.Pow(5, 2); v == w {
		fmt.Println("result is ture")
	} else {
		fmt.Println("error")
	}
}
