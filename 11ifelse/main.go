package main

import "fmt"

func main() {

	age := -1

	if age > 20 {
		fmt.Println("Adult")
	} else if age < 0 {
		fmt.Println("invalid")
	} else {
		fmt.Println("child")
	}

}
