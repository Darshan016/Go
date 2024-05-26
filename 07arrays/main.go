package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "luffy"
	names[1] = "light"
	names[2] = "torfinn"
	names[3] = "hiro"

	fmt.Println(names)
	fmt.Println("length is: ",len(names))


	var shows=[5]string{"oregairu","vinland saga","AOT","naruto"}
	fmt.Println(shows)
	fmt.Println("lenght is: ",len(shows))

}
