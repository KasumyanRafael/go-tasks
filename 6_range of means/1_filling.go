package main

import (
	"fmt"
	"math"
)

func main() {
	var countUint8Overflow, countUint16Overflow int

	for i := 0; i <= math.MaxUint32; i++ {
		if uint8(i) != i {
			countUint8Overflow++
		}
		if uint16(i) != i {
			countUint16Overflow++
		}
	}

	fmt.Printf("Количество переполнений для типа uint8: %d\n", countUint8Overflow)
	fmt.Printf("Количество переполнений для типа uint16: %d\n", countUint16Overflow)
}
