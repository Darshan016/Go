package main

import "fmt"

func main() {

	age := 22
	var myAge = &age

	fmt.Println(myAge)
	fmt.Println(*myAge)

	*myAge = *myAge + 100
	fmt.Println(*myAge)

}
