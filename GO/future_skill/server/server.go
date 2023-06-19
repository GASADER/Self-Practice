package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintf(w, "Hello World! %s \n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "Hello World! %s \n", method)
}

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe("localhost:3000", nil)
	fmt.Println("server running on localhost:3000")
	log.Fatal(err)
}
