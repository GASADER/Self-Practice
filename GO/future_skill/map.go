package main

import (
	"fmt"
	"strings"
)

func wordSplice(s string) map[string]int {
	//use build-in method
	words := strings.Fields(s)
	//like a empty array in js
	r := map[string]int{}
	//count some word
	for _, w := range words {
		r[w] = r[w] + 1
	}
	return r
}

func main() {
	//empty maps
	var a map[string]int
	fmt.Printf("%#v \n", a) //nil
	a = make(map[string]int)
	fmt.Printf("%#v \n", a) //{}

	//gat maps with default key and value
	var m map[string]int = map[string]int{}
	fmt.Println("value:", m)

	//add key and value
	m["result"] = 45
	fmt.Println("value:", m)
	m["answer"] = 42
	fmt.Println("value:", m)
	m["answer"] = 50
	fmt.Println("value:", m)

	//read value in key
	v := m["result"]
	fmt.Println("value:", v)

	//delete key and value
	delete(m, "answer")
	fmt.Println("value:", m)

	//check condition
	n, ok := m["answer"]
	if ok {
		fmt.Println("yes")
		fmt.Println("value:", n)
	} else {
		fmt.Println("no")
	}

	//cut string
	s := "Catch me if you can"
	w := wordSplice(s)
	fmt.Printf("%#v \n", w)
}
