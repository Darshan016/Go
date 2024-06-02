package main

import "fmt"

func main() {
	fmt.Println("One")
	newDefer()
	defer fmt.Println("Two")
	defer fmt.Println("three")
}

func newDefer(){
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}