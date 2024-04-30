package main

import "fmt"

func main() {
	var iq1, iq2, iq3 int

	fmt.Println("Введите IQ первого космонавта:")
	fmt.Scanln(&iq1)

	fmt.Println("Введите IQ второго космонавта:")
	fmt.Scanln(&iq2)

	fmt.Println("Введите IQ третьего космонавта:")
	fmt.Scanln(&iq3)

	var commander int

	if iq1 > iq2 {
		if iq1 > iq3 {
			commander = iq1
		} else {
			commander = iq3
		}
	} else {
		if iq2 > iq3 {
			commander = iq2
		} else {
			commander = iq3
		}
	}

	fmt.Println("Командир - космонавт с IQ:", commander)
}
