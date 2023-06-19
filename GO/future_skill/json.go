package main

import (
	"encoding/json"
	"fmt"
)

// make public struct and struct tag
type Coures struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func main() {
	c := Coures{
		Name:       "Basic go",
		Level:      "Basic",
		Views:      7562,
		Instructor: "A",
		FullPrice:  888,
	}

	//convert string to json with Marshal
	b, err := json.Marshal(c)
	fmt.Printf("type: %T \n", b)
	fmt.Printf("byte: %v \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Println(err)

	//get json data
	data := []byte(`{
			"name": "Basic go",
			"level": "Basic",
			"views": 7562,
			"instructor": "A",
			"full_price": 888
		}`)

	//convert json to string with Marshal
	var d Coures
	err2 := json.Unmarshal(data, &d)
	fmt.Printf("%#v \n", d)
	fmt.Println(err2)
}
