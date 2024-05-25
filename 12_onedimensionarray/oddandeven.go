package main

import (
	"fmt"
)

func main() {
	var number, evenCount, oddCount int
	for i := 0; i < 10; i++ {
		fmt.Print("Введите число: ")
		fmt.Scanln(&number)
		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Printf("Четных чисел: %d, Нечетных чисел: %d\n", evenCount, oddCount)
}
