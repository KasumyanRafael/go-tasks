package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			for mid > 0 && arr[mid-1] == target {
				mid--
			}
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{3, 8, 12, 15, 23, 31, 35, 44, 51, 55, 61, 72}

	fmt.Println("Упорядоченный массив:", arr)

	var num int
	fmt.Print("Введите число для поиска: ")
	fmt.Scanln(&num)

	index := binarySearch(arr, num)

	if index != -1 {
		fmt.Printf("Первое вхождение числа %d в массиве находится на позиции %d\n", num, index)
	} else {
		fmt.Println("Число не найдено в массиве.")
	}
}
