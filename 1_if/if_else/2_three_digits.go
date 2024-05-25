package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&num2)

	fmt.Println("Введите третье число:")
	fmt.Scanln(&num3)

	if num1 > 5 || num2 > 5 || num3 > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше 5.")
	}
}
