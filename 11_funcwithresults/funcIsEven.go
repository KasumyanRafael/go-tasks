package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0 //true or false
}

func main() {
	num := 5

	if isEven(num) {
		fmt.Printf("%d - число четное\n", num)
	} else {
		fmt.Printf("%d - число нечетное\n", num)
	}
}
