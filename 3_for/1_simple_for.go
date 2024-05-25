package main

import (
	"fmt"
)

func main() {
	var digit int
	fmt.Println("Введите число:")
	fmt.Scanln(&digit)

	for i := 0; i <= digit; i++ {
		println(i)
	}
}
