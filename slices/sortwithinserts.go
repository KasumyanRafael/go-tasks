package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	arr := []int{9, 5, 7, 2, 4, 1, 8, 3, 6, 0}

	fmt.Println("Массив до сортировки вставками:", arr)

	insertionSort(arr)

	fmt.Println("Массив после сортировки вставками:", arr)
}
