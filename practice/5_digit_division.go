package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&num2)

	if num2 != 0 {
		if num1%num2 == 0 {
			fmt.Println(num1, "делится на", num2, "без остатка.")
		} else {
			fmt.Println(num1, "не делится на", num2, "без остатка.")
		}
	} else {
		fmt.Println("На ноль делить нельзя!")
	}
}
