package main

import "fmt"

type course struct {
	name       string
	instructor string
	price      float64
}

//receiver
func (this course) discount(n float64) float64 {
	p := this.price - n
	return p
}
func (this course) info() {
	fmt.Println("name", this.name)
	fmt.Println("instructor", this.instructor)
	fmt.Println("price", this.price)
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

	d1 := c1.discount(20)
	d2 := c2.discount(20)
	d3 := c3.discount(20)
	fmt.Printf("discount: %v \n", d1)
	fmt.Printf("discount: %v \n", d2)
	fmt.Printf("discount: %v \n", d3)
	c1.info()
	c2.info()
	c3.info()
}
