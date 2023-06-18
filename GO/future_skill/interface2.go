package main

import "fmt"

//get naming of interface
type promotioner interface {
	discount() int
}
type infoer interface {
	info()
}

//Embedding interface
type presenter interface {
	promotioner
	infoer
}

//functions with interface
func sale(val promotioner) {
	fmt.Println("sale:", val.discount())
}
func show(val infoer) {
	val.info()
}
func summary(val presenter) {
	fmt.Println("sale:", val.discount())
	val.info()
}

type course struct{}

func (c course) discount() int {
	return 99
}

func (c course) info() {
	fmt.Println("c:", c)
}

func main() {
	v := course{}
	show(v)
	sale(v)
	summary(v)
}
