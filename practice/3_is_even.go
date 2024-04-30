package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Введите число:")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Число", number, "четное.")
	} else {
		fmt.Println("Число", number, "не четное.")
	}
}
