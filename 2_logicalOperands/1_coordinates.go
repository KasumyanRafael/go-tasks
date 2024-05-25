package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Введите значение X:")
	fmt.Scanln(&x)
	fmt.Println("Введите значение Y:")
	fmt.Scanln(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка находится в первой координатной четверти")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится во второй координатной четверти")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка находится в третьей координатной четверти")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в четвертой координатной четверти")
	} else {
		fmt.Println("Точка находится на одной из осей координат")
	}
}
