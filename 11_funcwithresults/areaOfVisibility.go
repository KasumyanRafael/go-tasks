package main

import "fmt"

var globalVar1 float64 = 5
var globalVar2 float64 = 10
var globalVar3 float64 = 15

func addGlobalVar1(input float64) float64 {
	result := input + globalVar1
	return result
}

func addGlobalVar2(input float64) float64 {
	result := input + globalVar2
	return result
}

func addGlobalVar3(input float64) float64 {
	result := input + globalVar3
	return result
}

func main() {
	input := 3.0

	result1 := addGlobalVar1(input)
	fmt.Printf("Результат после добавления глобальной переменной 1: %.2f\n", result1)

	result2 := addGlobalVar2(result1)
	fmt.Printf("Результат после добавления глобальной переменной 2: %.2f\n", result2)

	result3 := addGlobalVar3(result2)
	fmt.Printf("Результат после добавления глобальной переменной 3: %.2f\n", result3)
}
