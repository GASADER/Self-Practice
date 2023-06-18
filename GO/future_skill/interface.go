package main

import (
	"fmt"
)

// func show1(val int) {
// 	fmt.Println(val)
// }

// functions any type
func show2(val any) {
	//assign type
	// i, ok := val.(int)
	// if ok {
	// 	i = i + 2
	// 	fmt.Println(i)
	// } else {
	// 	fmt.Println("not int")
	// }

	// j, k := val.(string)
	// if k {
	// 	j = j + "hello"
	// 	fmt.Println(j)
	// } else {
	// 	fmt.Println("not string")
	// }

	switch v := val.(type) {
	case int:
		i := v + 2
		fmt.Println(i)
	case string:
		i := v + "hello"
		fmt.Println(i)
	default:
		fmt.Println("not handle type")
	}
}

func main() {
	// empty interface
	// var v interface{}

	// ver. 1.18 empty interface = any from Generics
	var v any
	v = 36
	fmt.Printf("type:%T value: %v \n", v, v)
	v = "hello"
	fmt.Printf("type:%T value: %v \n", v, v)

	//assign type
	// show1(v.(int))
	show2(v)
}
