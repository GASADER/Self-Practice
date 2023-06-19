package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

		//anonymous functions his looking to global scoop
		// defer func() {
		// 	// i = 10
		// 	fmt.Println(i)
		// }()
		//anonymous functions with parameter from variable local scoop
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

	fmt.Println("done")
}
