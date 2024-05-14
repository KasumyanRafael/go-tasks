package main

import (
	"fmt"
)

func calculateS(x int16, y uint8, z float32) float32 {
	result := 2*float32(x) + float32(y)*float32(y) - 3/z
	return result
}

func main() {
	x := int16(10)
	y := uint8(5)
	z := float32(2.5)

	S := calculateS(x, y, z)

	fmt.Printf("Результат вычислений: %0.2f\n", S)
}
