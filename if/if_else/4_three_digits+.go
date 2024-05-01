package main

import "fmt"

func main() {
	var num1, num2, num3, count int

	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&num2)

	fmt.Println("Введите третье число:")
	fmt.Scanln(&num3)

	if num1 >= 5 {
		count++
	}
	if num2 >= 5 {
		count++
	}
	if num3 >= 5 {
		count++
	}

	fmt.Println("Среди введённых чисел", count, "больше или равны 5.")
}
