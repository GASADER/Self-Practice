package main

import (
	"errors"
	"fmt"
)

//build-in interface error
// type error interface{
// 	Error() string
// }

// custom error
type myError struct{}

func (e myError) Error() string {
	return "myError"
}

// make new error from package error
var myerr = errors.New("my custom error")

func Divide(a, b float64) (float64, error) {
	//error handlers
	if b == 0 {
		err := myError{}
		return 0, err // fmt.Errorf("Divide by zero")
	}
	r := a / b
	//return error = nil
	return r, nil
}

func main() {
	//get variable for error
	r, err := Divide(1, 0)
	if err != nil {
		fmt.Println("handler error:", err)
		return
	}
	fmt.Println(r, err)
}
