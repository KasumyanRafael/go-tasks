package main

import "fmt"

func splitByEvenOdd(arr []int) ([]int, []int) {
	var even []int
	var odd []int

	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return even, odd
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even, odd := splitByEvenOdd(numbers)

	fmt.Println("Четные числа:", even)
	fmt.Println("Нечетные числа:", odd)
}
