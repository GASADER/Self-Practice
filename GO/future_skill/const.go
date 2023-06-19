package main

import "fmt"

// make custom type
type day int

//Varidic functions
func skills(xs ...string) {
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])

	//parameter is slices, can use method and property of slices
	for _, s := range xs {
		fmt.Println("skill:", s)
	}
}

func main() {
	//const is variable can't change value
	const (
		//iota his run value form 0 to length off variable
		_ = iota
		//get custom type for run type to variable
		sunday day = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	skills("js", "go", "py", "java")
}
