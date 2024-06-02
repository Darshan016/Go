package main

import "fmt"

func main() {

	characters := []string{"thorffin", "hiro", "lelouch", "L", "hikigaya", "itachi"}

	for i := 0; i < len(characters); i++ {
		// fmt.Println(characters[i])
	}

	for j := range characters {

		if characters[j] == "L" {
			break
		}

		fmt.Println(characters[j])
	}

}
