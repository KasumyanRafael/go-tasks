package main

import (
	"fmt"
)

func main() {
	var num1, num2 int16

	fmt.Print("Введите первое число (int16): ")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число (int16): ")
	fmt.Scanln(&num2)

	result := int32(num1) * int32(num2)

	var resultType string
	if result < -128 || result > 127 {
		resultType = "int8"
	} else if result < -32768 || result > 32767 {
		resultType = "int16"
	} else {
		resultType = "int32"
	}

	fmt.Printf("Результат умножения можно сохранить в типе данных: %s\n", resultType)
}
