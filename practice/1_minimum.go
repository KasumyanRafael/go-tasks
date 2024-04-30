package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&num2)

	var min int
	if num1 <= num2 {
		min = num1
	} else {
		min = num2
	}

	fmt.Println("Минимум из чисел", num1, "и", num2, "равен", min)
}
