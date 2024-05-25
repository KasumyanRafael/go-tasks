package main

import "fmt"

func main() {
	var first, second, third int

	for first < 10 || second < 100 || third < 1000 {
		if first < 10 {
			first++
		}
		if second < 100 {
			second++
		}
		if third < 1000 {
			third++
		}
	}
	fmt.Println("Значения переменных после цикла:")
	fmt.Println("Первая переменная:", first)
	fmt.Println("Вторая переменная:", second)
	fmt.Println("Третья переменная:", third)
}
