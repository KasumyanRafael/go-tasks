package main

import "fmt"

func swapValues(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 10, 20

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	swapValues(&a, &b)

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
