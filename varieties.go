package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 153
	fmt.Println("До обмена значений местами")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// обмен местами значений переменных

	fmt.Println("После обмена значений местами")
	fmt.Println("a:", b)
	fmt.Println("b:", a)
}
