package main

import "fmt"

func main() {
	greet()
	ans := getSum(5, 5)
	fmt.Println(ans)

	sum := numSum(2, 2, 2,2,2)
	fmt.Println(sum)
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func numSum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func greet(){
	fmt.Println("Hey there")
}
