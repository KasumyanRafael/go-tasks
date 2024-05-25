package main

import (
	"fmt"
)

func determinant(matrix [3][3]int) int {
	det := 0

	// Расчет определителя по формуле Sarrus
	a := matrix[0][0] * matrix[1][1] * matrix[2][2]
	b := matrix[0][1] * matrix[1][2] * matrix[2][0]
	c := matrix[0][2] * matrix[1][0] * matrix[2][1]
	d := matrix[0][2] * matrix[1][1] * matrix[2][0]
	e := matrix[0][0] * matrix[1][2] * matrix[2][1]
	f := matrix[0][1] * matrix[1][0] * matrix[2][2]

	det = a + b + c - d - e - f

	return det
}

func main() {
	matrix := [3][3]int{
		{2, 5, 1},
		{6, 3, 4},
		{5, 2, 7},
	}

	fmt.Println("Исходная матрица:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	det := determinant(matrix)

	fmt.Println("Определитель матрицы:", det)
}
