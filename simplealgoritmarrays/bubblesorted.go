package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	array := []int{6, 5, 3, 1, 8, 2}

	fmt.Println("Исходный массив:", array)

	bubbleSort(array)

	fmt.Println("Отсортированный массив:", array)
}
