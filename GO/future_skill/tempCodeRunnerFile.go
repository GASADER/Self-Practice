package main

import "fmt"

type promotion interface {
	discount() int
}

func sale(val promotion) {
	fmt.Println("sale:", val)
}

type course struct{}

func (c course) discount() int {
	return 99
}

func main() {
	v := course{}
	sale(v)
}
