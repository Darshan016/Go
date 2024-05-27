package main

import "fmt"

func main() {

	var animes = []string{"bleach", "gundam", "naruto"}
	fmt.Println(animes)
	fmt.Printf("Type is: %T\n", animes)

	animes = append(animes, "DITF")
	fmt.Println(animes)

	heroes := make([]string, 4)
	heroes[0] = "A"
	heroes[1] = "B"
	heroes[2] = "C"
	heroes[3] = "D"
	fmt.Println(heroes)
	heroes = append(heroes,"E")
	fmt.Println(heroes)


	//removing an item from a slice

	var index int = 2
	heroes = append(heroes[:index],heroes[index+1:]...)
	fmt.Println(heroes)

}
