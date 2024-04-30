package main

import (
	"fmt"
)

func main() {
	var expPoints int
	fmt.Println("Введите количество очков опыта:")
	fmt.Scanln(&expPoints)

	level := 1
	if expPoints >= 1000 {
		level = 2
	}
	if expPoints >= 2500 {
		level = 3
	}
	if expPoints >= 5000 {
		level = 4
	}

	fmt.Println("Ваш уровень:", level)
}
