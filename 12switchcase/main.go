package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to ludo")

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6) + 1

	switch number {
	case 1:
		fmt.Println("start your move.")
	case 2:
		fmt.Println("move 2 places")
	case 3:
		fmt.Println("move 3 places")
		fallthrough
	case 4:
		fmt.Println("move 4 places")
	case 5:
		fmt.Println("move 5 places")
	case 6:
		fmt.Println("move 6 places and roll the dice again")
	default:
		fmt.Println("invalid number")
	}

}
