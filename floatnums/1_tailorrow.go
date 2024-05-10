package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var precision int

	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	fmt.Print("Введите требуемую точность после запятой: ")
	fmt.Scan(&precision)

	result := 0.0
	factorial := 1.0
	n := 0

	for {
		term := math.Pow(x, float64(n)) / factorial
		result += term
		n++
		factorial *= float64(n)

		if math.Abs(term) < math.Pow(10, float64(-precision)) {
			break
		}
	}

	fmt.Printf("Результат вычисления e^%v с точностью до %d знаков после запятой: %.*f\n", x, precision, precision, result)
}
