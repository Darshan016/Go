package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println(today)

	newd := today.Format("2006-01-02")
	fmt.Println(newd)

	fmt.Println(today.Format("sunday, jan 2, 2001"))
}
