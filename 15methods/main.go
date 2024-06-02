package main

import "fmt"

type Show struct {
	Name        string
	FemaleLead  string
	MaleLead    string
	IsCompleted bool
}

func main() {
	DITF := Show{"Darling in the franxx", "zero two", "hiro", true}
	fmt.Printf("show details are: %+v\n", DITF)

	fmt.Println(DITF)
	fmt.Println(DITF.FemaleLead)
	fmt.Println(DITF.MaleLead)
	getMaleLead(DITF)
}

func getMaleLead(s Show) {
	fmt.Println("MC is: ", s.MaleLead)
}
