package main

import "fmt"

func show(sk [3]string) {
	fmt.Printf("result = %#v\n", sk)
}

func main() {
	// empty array
	// var skills [3]string
	// array with default value
	var skills [3]string = [3]string{"js", "go", "c#"}
	fmt.Printf("%#v\n", skills)

	//get value in array
	v := skills[1]
	fmt.Printf("v:%#v \n", v)

	//find length of array
	l := len(skills)
	fmt.Printf("len: %#v \n", l)

	//add value in array
	skills[1] = "golang"
	fmt.Printf("%#v \n", skills)
	show(skills)

	// var item [2]string
	// show(item)
}
