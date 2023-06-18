package main

import (
	"fmt"
)

func main() {
	sum := 1
	// for loop
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println("i:", i, "sum:", sum)
	}

	//while loop
	for sum < 100 {
		sum += sum
		fmt.Println("sum:", sum)
	}

	//infinity loop
	for {
		fmt.Println("infinity")
		break
	}

	//loop in array
	skills := [3]string{"go", "js", "c#"}
	for i := 0; i < len(skills); i++ {
		fmt.Println(skills[i])
	}

	//for range
	for i := range skills {
		fmt.Println(skills[i])
	}

	for i, val := range skills {
		fmt.Println("index:", i, "value:", val)
	}

	for _, val := range skills {
		fmt.Println("value:", val)
	}
}
