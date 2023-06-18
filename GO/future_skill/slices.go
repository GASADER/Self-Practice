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

	//get variable & add multivalue
	ss := append(skills, "kotlin", "dart")
	fmt.Printf("length: %d -value: %#v \n", len(ss), ss)

	//Half-open range ,Slice of Slices
	s1 := skills[0:2]
	s2 := skills[1:3]
	s3 := skills[:]
	fmt.Printf("length: %d -value: %#v \n", len(s1), s1)
	fmt.Printf("length: %d -value: %#v \n", len(s2), s2)
	fmt.Printf("length: %d -value: %#v \n", len(s3), s3)

	//capacity
	s4 := skills[0:3]
	c4 := cap(s4)
	fmt.Printf("length: %d -capacity:%#v -value: %#v \n", len(s4), c4, s4) //length: 3 -capacity:6 -value: []string{"js", "go", "c#"}
	//capacity:6 form  skills 3 , skills append +1, ss append +2

	//underlying array
	s5 := skills[1:4]
	c5 := cap(s5)
	fmt.Printf("length: %d -capacity:%#v -value: %#v \n", len(s5), c5, s5) //length: 3 -capacity:5 -value: []string{"js", "go", "c#", "ruby"}
	//capacity:5 form  skills 3 , skills append +1, ss append +2, start splice -1

	//return capacity for append
	s6 := skills[1:6]
	s6 = append(s6, "java")
	c6 := cap(s6)
	fmt.Printf("length: %d -capacity:%#v -value: %#v \n", len(s6), c6, s6) //length: 6 -capacity:10 -value: []string{"go", "c#", "ruby", "kotlin", "dart", "java"}
	//capacity:10 form  skills 3 , skills append +1, ss append +2, s6 append +1 and new generate +3

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

	//ellipsis operator (spread operator)
	arr1 := []int{0, 1, 2, 3}
	arr2 := []int{4, 5, 6, 7}

	var mix []int
	mix = append(arr1, arr2...)
	fmt.Printf("length: %d -value: %#v \n", len(mix), mix)

}
