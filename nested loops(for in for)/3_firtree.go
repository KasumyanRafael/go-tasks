package main

import "fmt"

func main() {
	var height int

	fmt.Print("Введите высоту вашей ёлочки: ")
	fmt.Scanln(&height)
	for i := 1; i <= height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
