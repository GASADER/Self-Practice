package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	c1 := course{
		name:       "Basic Go",
		instructor: "A",
		price:      100,
	}
	var c2 = course{"Go", "C", 200}
	var c3 = course{instructor: "C"}

	name := c1.name
	//assign statement
	c1.instructor = "B"
	fmt.Println("name:", name)
	fmt.Println("instructor:", c1.instructor)
	fmt.Println("price:", c1.price)

	fmt.Println("name:", c2.name)
	fmt.Println("instructor:", c2.instructor)
	fmt.Println("price:", c2.price)

	fmt.Println("c3:", c3)
}
