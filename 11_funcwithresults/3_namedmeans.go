package main

import (
	"fmt"
)

func multiplyNumber(num float64) (result float64) {
	result = num * 2
	return
}

func addNumber(num float64) (result float64) {
	result = num + 10
	return
}

func main() {
	var input float64
	fmt.Print("Введите число: ")
	fmt.Scanln(&input)

	multiplied := multiplyNumber(input)
	added := addNumber(multiplied)

	fmt.Printf("Введенное число: %.2f\n", input)
	fmt.Printf("После умножения: %.2f\n", multiplied)
	fmt.Printf("После прибавления: %.2f\n", added)
}
