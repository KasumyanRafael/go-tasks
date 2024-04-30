package main

import "fmt"

func main() {
	var year int
	fmt.Println("Введите год:")
	fmt.Scanln(&year)

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("Год является високосным. В нем 366 дней.")
	} else {
		fmt.Println("Год не является високосным. В нем 365 дней.")
	}
}
