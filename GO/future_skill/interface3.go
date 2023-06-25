package main

import "fmt"

//
type Animal interface {
	MakeSound()
	Move()
}

type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("Woof!")
}

func (d Dog) Move() {
	fmt.Println("Running")
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("Meow!")
}

func (c Cat) Move() {
	fmt.Println("Jumping")
}

func main() {
	var animal Animal

	animal = Dog{}
	animal.MakeSound() // Output: Woof!
	animal.Move()      // Output: Running

	animal = Cat{}
	animal.MakeSound() // Output: Meow!
	animal.Move()      // Output: Jumping
}
