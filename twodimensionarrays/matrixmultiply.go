package main

import (
	"fmt"
)

func multiplyMatrices(matrix1 [3][5]int, matrix2 [5][4]int) [3][4]int {
	result := [3][4]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 5; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

func main() {
	matrix1 := [3][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	matrix2 := [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	result := multiplyMatrices(matrix1, matrix2)

	fmt.Println("Результирующая матрица:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}
