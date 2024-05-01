package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Введите два числа:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	sum := num1
	for sum < num1+num2 {
		sum++
	}
	fmt.Println("Сумма:", sum)
}
