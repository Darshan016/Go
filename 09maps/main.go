package main

import "fmt"

func main() {
	shows := make(map[string]string)
	shows["DITF"] = "Darling in the franxx"
	shows["DN"] = "Death Note"
	fmt.Println(shows)


	for k,v := range(shows){
		fmt.Println(k,"=",v)
	}

	delete(shows,"DN")
	fmt.Println(shows)
}
