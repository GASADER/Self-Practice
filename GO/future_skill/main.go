package main

import (
	"fmt"
)

// void functions
func add1() {
	fmt.Println("test")
}

// parameter funtions & return value
func add2(x int, y int) int {
	fmt.Printf("%d and %d", x, y)
	result := x + y
	return result
}

// return mutiple declaration
func add3(x, y int) (int, string) {
	number := x + y
	word := "hello"
	fmt.Println("result:", x, y)
	return number, word
}

// swap function
func swap(x, y int) (int, int) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum % 2
	return
}

// first class functions
var i string = "hello"
var j int = 4
var first = func(i, j float64) float64 {
	return i + j
}

// Higher-Order Functions (HOF)
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

//Higher-Order Functions (HOF) Multiple functions parameter

func inc() int {
	return 1
}

func curr(x int) int {
	x = 2
	return x
}

func adder() (func() int, func(x int) int) {
	return inc, curr
}

// Higher-Order Function (HOF) with Recursive functions
func applyRecursively(n int, f func(int) int) int {
	if n == 0 {
		return 0
	}
	return f(n) + applyRecursively(n-1, f)
}

func square(n int) int {
	return n * n
}

// run functions
func main() {
	add1()

	a := add2(50, 100)
	fmt.Println(a)

	b, c := add3(50, 100)
	fmt.Println(b, c)
	d, e := swap(50, 100)
	fmt.Println(d, e)

	f, g := split(50)
	fmt.Println(f, g)

	h := first(50, 100)
	fmt.Println(h)
	fmt.Printf("%T \n", first) //func(float64, float64) float64

	result := applyOperation(5, 3, add)
	fmt.Println(result)
	result = applyOperation(5, 3, multiply)
	fmt.Println(result)

	i, k := adder()
	fmt.Println(i, k) //value is address of functions in memory  //0x57fa60 0x57fa80
	fmt.Println(i(), k(2))

	n := 5
	sum := applyRecursively(n, square)
	fmt.Println("Sum of squares from 1 to", n, "is", sum)
}
