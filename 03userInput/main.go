package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")

	input, _ := reader.ReadString('\n')
	fmt.Println("User age is ", input)
	fmt.Printf("type of age is %T \n", input)

}
