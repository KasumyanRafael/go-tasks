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

	// Проверка наличия хотя бы одного положительного числа
	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Хотя бы одно число является положительным")
	} else {
		fmt.Println("Все введенные числа отрицательны или равны нулю")
	}
}
