package main

import "fmt"

func main() {
	skills := []string{"js", "go", "c#"}
	fmt.Printf("%#v\n", skills)

	//zero value
	var zero1 []int
	fmt.Printf("length: %d -value: %#v \n", zero1) //length: [] -value: %!v(MISSING)

	//index out of range
	// z := zero[0]
	//zero[0] = 33
	// fmt.Printf("length: %d -value: %#v \n", z)

	//make
	zero2 := make([]int, 3)
	fmt.Printf("length: %d -value: %#v \n", zero2)

	//append for increase
	skills = append(skills, "ruby")
	fmt.Printf("length: %d -value: %#v \n", len(skills), skills)

	//get variable & add mutivalue
	ss := append(skills, "kotlin", "dart")
	fmt.Printf("length: %d -value: %#v \n", len(ss), ss)

	//Half-open range ,Slice of Slices
	s1 := skills[0:2]
	s2 := skills[1:3]
	s3 := skills[:]
	fmt.Printf("length: %d -value: %#v \n", len(s1), s1)
	fmt.Printf("length: %d -value: %#v \n", len(s2), s2)
	fmt.Printf("length: %d -value: %#v \n", len(s3), s3)

	// convert array to slices
	arr := [...]int{0, 1, 2, 3, 4}
	slices := arr[:]
	fmt.Printf("length: %d -value: %#v \n", len(slices), slices)

	//convert slices to array
	slice := []int{1, 2, 3, 4, 5}
	array := [5]int{}

	for index, value := range slice {
		array[index] = value
	}

	fmt.Printf("length: %d -value: %#v \n", len(slice), slice)

	//copy for decrease
	decrease := []int{1, 2, 3, 4, 5}
	newSlice := make([]int, len(decrease)-2)
	copy(newSlice, decrease[2:])
	fmt.Printf("length: %d -value: %#v \n", len(newSlice), newSlice)
}
