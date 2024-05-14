package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 5, 7, 2, 4, 1, 8, 3, 6, 0}

	// Анонимная функция для сортировки пузырьком и переворачивания массива
	sortAndReverse := func(array []int) []int {
		n := len(array)
		swapped := true

		for swapped {
			swapped = false
			for i := 1; i < n; i++ {
				if array[i-1] > array[i] {
					array[i-1], array[i] = array[i], array[i-1]
					swapped = true
				}
			}
		}

		// Переворачивание массива
		for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
			array[i], array[j] = array[j], array[i]
		}

		return array
	}

	fmt.Println("Исходный массив:", arr)

	// Вызов анонимной функции и получение отсортированного и перевернутого массива
	sortedReversed := sortAndReverse(arr)
	fmt.Println("Отсортированный и перевернутый массив:", sortedReversed)
}
