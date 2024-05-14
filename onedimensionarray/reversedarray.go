package main

import "fmt"

func reverseArray(arr []int) []int {
	reversed := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}
	return reversed
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Исходный массив:", array)

	reversedArray := reverseArray(array)

	fmt.Println("Инвертированный массив:", reversedArray)
}
