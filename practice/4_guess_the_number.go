package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 3 // Начнем с середины диапазона

	fmt.Println("Загадайте число от 1 до 5")

	for {
		fmt.Printf("Это число %d? (да/меньше/больше): ", guess)
		var answer string
		fmt.Scanln(&answer)

		switch answer {
		case "да":
			fmt.Println("Ура! Я угадал!")
			return
		case "меньше":
			guess -= 1
		case "больше":
			guess += 1
		default:
			fmt.Println("Пожалуйста, введите одно из допустимых значений: да, меньше или больше.")
		}
	}
}
