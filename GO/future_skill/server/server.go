package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	//variable method
	method := r.Method
	fmt.Fprintf(w, "Hello World! %s \n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "Hello World! %s \n", method)

	//check condition from method
	if method == "post" {
		//
		ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "%s", r)
		return
	}
}

func main() {
	//run functions from path directory
	http.HandleFunc("/", greet)

	//run server
	err := http.ListenAndServe("localhost:3000", nil)
	fmt.Println("server running on localhost:3000")
	log.Fatal(err)
}
