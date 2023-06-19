package main

import "fmt"

//Dereference
func changePricePointer(p *int) {
	*p = *p - 599
	fmt.Println("change", p, &p)
}

func changePrice(price int) {
	price = price - 599
	fmt.Println("change", price, &price)
}

func main() {
	var price int = 9999
	var addr *int = &price

	fmt.Println(price, addr)

	//write new value in pointer
	*addr = 9400
	fmt.Println(price, addr)

	//read value in pointer
	v := *addr
	fmt.Println(v)

	//arithmetic pointer
	//not working in go
	// addr ++

	//get value form variable to parameter
	changePricePointer(&price)
	fmt.Println(price, addr)

	changePrice(price)
	fmt.Println(price, addr)

}
