package main

import (
	"encoding/json"
	"fmt"
)

type Coures struct {
	Name       string
	Level      string
	Views      int
	Instructor string
	FullPrice  int
}

func main() {
	c := Coures{
		Name:       "Basic go",
		Level:      "Basic",
		Views:      7562,
		Instructor: "A",
		FullPrice:  888,
	}

	b, err := json.Marshal(c)
	fmt.Printf("type: %T \n", b)
	fmt.Printf("byte: %v \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Println(err)
}
