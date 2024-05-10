package main

import "fmt"

func sumOfEvenNumbersInRange(start, end int) {
	sum := 0
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("Сумма чётных чисел в диапазоне от %d до %d составляет: %d\n", start, end, sum)
}

func main() {
	start := 1
	end := 10

	sumOfEvenNumbersInRange(start, end)
}
