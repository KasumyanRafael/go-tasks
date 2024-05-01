package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	// Запрос пользователю ввести три числа
	fmt.Println("Введите три числа:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	// Проверка наличия совпадающих чисел
	if num1 == num2 || num1 == num3 || num2 == num3 {
		fmt.Println("Есть совпадающие числа")
	} else {
		fmt.Println("Совпадающих чисел нет")
	}
}
